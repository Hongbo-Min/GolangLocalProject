package geecache

import "Hongbo-Min-github/GolangLocalProject/7-days-Gee/geecache/geecachepb"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(in *geecachepb.Request, out *geecachepb.Response) error
}
