package archive_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/goulash/archive"
)

func TestReadFileFromArchive(z *testing.T) {
	const testdataPath = "testdata/dir_reader_data"
	const archivePath = "dir1/file1"
	const archiveData = "dir1/file1 content\n"

	_, err := archive.ReadFileFromArchive(testdataPath+".tar.db", archivePath)
	fmt.Printf("expected error: %v\n", err)
	if err == nil {
		z.Fatalf("expected error, got nil")
	}

	for _, ext := range []string{
		".tar", ".tar.gz", ".tar.bz2", ".tar.xz", ".tar.zst",
		".tar.gz.db", ".tar.bz2.db", ".tar.xz.db", ".tar.zst.db",
	} {
		path := testdataPath + ext
		fmt.Printf("read archive: %s\n", path)
		bs, err := archive.ReadFileFromArchive(path, archivePath)
		if err != nil {
			z.Errorf("%s: unexpected error: %v", path, err)
			continue
		}
		if string(bs) != archiveData {
			z.Errorf("%s: expected data %q, got %q", path, archiveData, bs)
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
