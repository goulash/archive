package archive_test

import (
	"testing"

	"github.com/goulash/archive"
)

func TestReadFileFromArchive(z *testing.T) {
	const testdataPath = "testdata/dir_reader_data"
	const archivePath = "dir1/file1"
	const archiveData = "dir1/file1 content\n"
	for _, ext := range []string{".tar", ".tar.gz", ".tar.bz2", ".tar.xz"} {
		bs, err := archive.ReadFileFromArchive(testdataPath+ext, archivePath)
		if err != nil {
			z.Fatalf("unexpected error: %s", err)
		}
		if string(bs) != archiveData {
			z.Errorf("expected data %q, got %q", archiveData, bs)
		}
	}
}
