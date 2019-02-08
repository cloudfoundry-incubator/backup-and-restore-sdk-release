package s3

const blobDelimiter = "/"

type Blob struct {
	path string
}

func NewBlob(path string) Blob {
	return Blob{
		path: path,
	}
}

func (b Blob) Path() string {
	return b.path
}