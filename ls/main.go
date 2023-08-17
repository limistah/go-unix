package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"text/tabwriter"
	"time"
)

func getHardLinksToPath(fpath string) (uint64, error) {
	fi, err := os.Stat(fpath)
	if err != nil {
		return 0, err
	}
	nlink := uint64(0)
	if sys := fi.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			nlink = uint64(stat.Nlink)
		}
	}
	return nlink, err
}

func main() {
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	files, err := os.ReadDir(basePath)
	if err != nil {
		panic(err)
	}
	files = append(files)
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 1, ' ', tabwriter.AlignRight)
	for _, p := range []string{".", ".."} {
		printFileInfo(w, p)
	}
	for _, f := range files {
		printFileInfo(w, f.Name())
	}
	w.Flush()
}

func printFileInfo(w *tabwriter.Writer, p string) {
	info, err := os.Stat(p)
	stat, _ := info.Sys().(*syscall.Stat_t)

	if err != nil {
		fmt.Println("err: unable to read path")
	}
	nlink, err := getHardLinksToPath(info.Name())
	if err != nil {
		fmt.Println("err: unable to get link to file")
	}

	grp, err := user.LookupGroupId(strconv.Itoa(int(stat.Gid)))
	if err != nil {
		fmt.Println("err: unable to get link to file")
	}
	usr, err := user.LookupId(strconv.Itoa(int(stat.Uid)))
	if err != nil {
		fmt.Println("err: unable to get link to file")
	}

	fmt.Fprint(w, info.Mode())
	fmt.Fprint(w, fmt.Sprintf("\t%d\t", nlink))
	fmt.Fprint(w, fmt.Sprintf("\t%s\t", usr.Username))
	fmt.Fprint(w, fmt.Sprintf("\t%s\t", grp.Name))
	fmt.Fprint(w, fmt.Sprintf("\t%d\t", info.Size()))
	yr := info.ModTime().Year()
	date := info.ModTime().Format("Jan _2 2006")
	if time.Now().Year() == yr {
		date = info.ModTime().Format("Jan _2 15:04")	
	}
	
	fmt.Fprint(w, fmt.Sprintf("\t%s\t", date))
	fmt.Fprint(w, fmt.Sprintf("\t%s\n", info.Name()))
}
