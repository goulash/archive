package archive_test

import (
	"os"
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

func TestExtractArchive(z *testing.T) {
	const testdataArchive = "testdata/repoctl.tar.gz"
	const destdir = "testdata/tmp"

	defer func() {
		os.RemoveAll(destdir)
	}()

	err := archive.ExtractArchive(testdataArchive, destdir)
	if err != nil {
		z.Errorf("unexpected error: %s", err)
	}
}
