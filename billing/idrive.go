package "idrive"


type XMLUser struct {
	XMLName xml.Name `xml:"name"`
	Key string `xml:"name,attr"`
}

type XMLStorageReport struct {
	XMLName xml.Name `xml:"StorageReport"`

}

