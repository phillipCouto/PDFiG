package main

import (
	"github.com/mawicks/PDFiG/pdf" )

// make_file() produces a file using low-level methods of the pdf.File
// type.  It does not work at the document layer and it does *not*
// produce a PDF document that a viewer will understand.
func modify_file() {
	f := pdf.OpenFile("/tmp/test-document.pdf")
	documentInfo := pdf.NewDocumentInfo()
	documentInfo.SetTitle("Rewritten Title")
	documentInfo.SetAuthor("Nobody")
	documentInfo.SetCreator("Nothing")

	documentInfoIndirect := pdf.NewIndirect(f)
	documentInfoIndirect.Finalize(documentInfo)
	f.SetInfo (documentInfoIndirect)

	f.Close()
}
