package image

import "testing"

const (
	defaultSize float64 = 224
)

func TestMD5Hash(t *testing.T) {
	tt := []struct {
		name     string
		path     string
		size     float64
		expected string
	}{
		{
			name:     "incorrect image path",
			path:     "AAAAAAAAAAAAAAAA",
			expected: "",
			size:     defaultSize,
		},
		{
			name:     "0120001c599dc0af44615e7fe8c41b255a0d288d",
			path:     "./testdata/0120001c599dc0af44615e7fe8c41b255a0d288d",
			expected: "4acb5e65380dc157a69f47c1bb9b5daa",
			size:     defaultSize,
		},
		{
			name:     "012000204363b62eafe245f624c0cf36518880bb",
			path:     "./testdata/012000204363b62eafe245f624c0cf36518880bb",
			expected: "a2a71b51fa1edc03bbf965a729eb7072",
			size:     defaultSize,
		},
		{
			name:     "01200024995b6fcca6a3f5ee275de74ddd93b4e6",
			path:     "./testdata/01200024995b6fcca6a3f5ee275de74ddd93b4e6",
			expected: "5ca7a0bf0d64af9aee76ad3fc0def6e8",
			size:     defaultSize,
		},
		{
			name:     "012000428fdc45ec7ad7e63cc3404627b0d54815",
			path:     "./testdata/012000428fdc45ec7ad7e63cc3404627b0d54815",
			expected: "cae9aa333c3cedaded117c1faef1e2c4",
			size:     defaultSize,
		},
		{
			name:     "01200048007a32201be85bc43aafd7de5526fa25",
			path:     "./testdata/01200048007a32201be85bc43aafd7de5526fa25",
			expected: "781fa5a025583535705bf277ab0b7c30",
			size:     defaultSize,
		},
		{
			name:     "0120004c1a13827989dd9ac19875c1f0719c775b",
			path:     "./testdata/0120004c1a13827989dd9ac19875c1f0719c775b",
			expected: "83dca2847371b919448992a186751aeb",
			size:     defaultSize,
		},
		{
			name:     "0120004f21edb7cfc3184ff572a85c4e5edde6b4",
			path:     "./testdata/0120004f21edb7cfc3184ff572a85c4e5edde6b4",
			expected: "11164c312daf66fb5f5dc1e43ef1d21b",
			size:     defaultSize,
		},
		{
			name:     "01200073904a56e3da52af6330c6f380fe235503",
			path:     "./testdata/01200073904a56e3da52af6330c6f380fe235503",
			expected: "af2aa16f686a0dda983cb706370cfa79",
			size:     defaultSize,
		},
		{
			name:     "01200074bbb9fb5ba793db0196479fcee8e36176",
			path:     "./testdata/01200074bbb9fb5ba793db0196479fcee8e36176",
			expected: "17ce4e5078f2e8f24f7a72f1e6b84945",
			size:     defaultSize,
		},
		{
			name:     "0120007f33c98b575290ba9487254cf73f0a6247",
			path:     "./testdata/0120007f33c98b575290ba9487254cf73f0a6247",
			expected: "b2740f1277685ca22bc4ad2a1351a2b3",
			size:     defaultSize,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := MD5Hash(tc.path, tc.size)
			if got != tc.expected {
				t.Fatalf("unexpected error: got=%s but expected=%s\n", got, tc.expected)
			}
		})
	}
}
