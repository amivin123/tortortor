package torrent

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/amivin123/tortortor/internal/testutil"
	"github.com/amivin123/tortortor/storage"
)

func TestHashPieceAfterStorageClosed(t *testing.T) {
	td, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(td)
	tt := &Torrent{
		storageOpener: storage.NewClient(storage.NewFile(td)),
	}
	mi := testutil.GreetingMetaInfo()
	info, err := mi.UnmarshalInfo()
	require.NoError(t, err)
	require.NoError(t, tt.setInfo(&info))
	require.NoError(t, tt.storage.Close())
	tt.hashPiece(0)
}
