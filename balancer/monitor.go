package balancer

import (
	"log"
	"os"
	"sync"

	lru "github.com/hashicorp/golang-lru"
	"github.com/mirrorhub-io/platform/client"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
)

type Service struct {
	Orig      *pb.Service
	Files     []*File
	FileCache *lru.ARCCache
	Monitor   *Monitor
}

const (
	FileOnline  = 1
	FileOffline = 2
)

type File struct {
	State  int
	Path   *string
	Mirror *pb.Mirror
}

const (
	FileCacheSize    = 1024
	ServiceCacheSize = 512
)

type Monitor struct {
	Client       *client.Client
	Logger       *log.Logger
	Services     []*Service
	ServiceCache *lru.ARCCache
	Mutex        *sync.Mutex
}

func NewMonitor(c *client.Client) *Monitor {
	l, _ := lru.NewARC(ServiceCacheSize)
	return &Monitor{
		Client:       c,
		Logger:       log.New(os.Stdout, " [Monitor] ", 0),
		ServiceCache: l,
		Mutex:        &sync.Mutex{},
	}
}

func (m *Monitor) Authorize() error {
	if _, err := m.Client.Contact().Authorize(); err != nil {
		m.Logger.Println("Authorize error. ", err)
		return err
	}
	return nil
}

func (m *Monitor) Preload() {
	m.Logger.Println("Preload monitor")
	if m.Authorize() != nil {
		return
	}
	list, _ := m.Client.Service().List()
	services := make([]*Service, len(list.Services))
	for i, service := range list.Services {
		services[i] = NewService(service, m)
	}
	m.Mutex.Lock()
	m.Logger.Println("Services:", len(services))
	m.Services = services
	m.Mutex.Unlock()
}

func NewService(s *pb.Service, m *Monitor) *Service {
	l, _ := lru.NewARC(FileCacheSize)
	ss := &Service{
		Orig:      s,
		FileCache: l,
		Monitor:   m,
	}
	ss.Preload()
	return ss
}

func (s *Service) Preload() {
	files := make([]*File, len(s.Orig.Files))
	for i, file := range s.Orig.Files {
		files[i] = &File{
			State: FileOffline,
			Path:  &file,
		}
	}
	s.Monitor.Logger.Println("Files:", len(files))
}
