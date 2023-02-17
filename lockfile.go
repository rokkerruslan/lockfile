package lockfile

import (
	"os"
	"path/filepath"
)

// todo: any profit of advisory locks (fcntl, flock) over opening a
//       separate file with O_CREATE|O_EXCL?
//       https://man7.org/linux/man-pages/man2/fcntl.2.html

// doc: Mandatory File Locking For The Linux Operating System
//      https://www.kernel.org/doc/Documentation/filesystems/mandatory-locking.txt

const lockfile = "LOCKFILE"
const lockfileMode = os.FileMode(0600)
const lockfileFlags = os.O_CREATE | os.O_EXCL

// Lock locks directory. All another Lock calls will be return an error.
func Lock(path string) error {
	f, err := os.OpenFile(filepath.Join(path, lockfile), lockfileFlags, lockfileMode)
	if err != nil {
		return err
	}

	return f.Close()
}

// Unlock unlocks directory.
func Unlock(path string) error {
	return os.Remove(filepath.Join(path, lockfile))
}
