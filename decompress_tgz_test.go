package getter

import (
	"path/filepath"
	"testing"
)

func TestTarGzipDecompressor(t *testing.T) {
	cases := []TestDecompressCase{
		{
			"empty.tar.gz",
			false,
			true,
			nil,
			"",
		},

		{
			"single.tar.gz",
			false,
			false,
			nil,
			"d3b07384d113edec49eaa6238ad5ff00",
		},

		{
			"single.tar.gz",
			true,
			false,
			[]string{"file"},
			"",
		},

		{
			"multiple.tar.gz",
			true,
			false,
			[]string{"file1", "file2"},
			"",
		},

		{
			"multiple.tar.gz",
			false,
			true,
			nil,
			"",
		},
	}

	for _, tc := range cases {
		tc.Input = filepath.Join("./test-fixtures", "decompress-tgz", tc.Input)
	}

	TestDecompressor(t, new(TarGzipDecompressor), cases)
}
