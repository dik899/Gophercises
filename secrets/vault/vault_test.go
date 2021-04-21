package secret

import(
	"testing"
	"path/filepath"
    homedir "github.com/mitchellh/go-homedir"
)
func TestLoad(t *testing.T) {
	var v Vault
v.load()
}

func TestSave(t *testing.T) {
var v Vault
 v.save()

}

func TestSet(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secret.txt")
	  v := File("abc", fp)
	  v.Set("xyz", "testing")
	
	
}
func TestSetNegative(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secret.txt")
	v := File("", fp)
	res := v.Set("xyz", "testing")
	if res == nil {
		t.Error("No Error")
	}
}

func TestGet(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secret.txt")
	v := File("abc", fp)
	
	v.Get("xyz")

}
func TestGetNegative(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secret.txt")
	v := File("abc", fp)
	_, err:= v.Get("x")
	if err == nil {
		t.Error("NO error ")
	}
	v = File("", fp)
	_, err = v.Get("x")
	if err == nil {
		t.Error(" NO error ")
	}
}