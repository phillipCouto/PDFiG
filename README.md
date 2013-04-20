Portable Document Format In Go (PDFiG)
======================================It is meant for low-level manipulation of PDF files and
will not necessarily produce a file that can be read by a PDF reader
without imposing additional document structure.

Read, write, and modify PDF files using the Go programming language

This project has the goal of being a full-featured PDF library with
the ability to read, writing, and modifying PDF files.  The
architecture has been designed to support having multiple PDF files
open and moving objects from one to another.  Object references read
from one PDF file are transparently translated and renumbered when
used in another PDF file.  The project has the goal (partially
implemented) of being able to read arbitrary objects from an open PDF
file and use them in another PDF file.  It has an additional goal of
being able to make append-only, incremental revisions to existing PDF
files.  The design of the Portable Document Format specifically
provides for such revisions.

The API currently has two major types: a `pdf.File` and a
`pdf.Document`.  The `pdf.File` type deals with low-level file
structure.  The `pdf.Document` type is a high-level API that uses an
internal `pdf.File` object to generate usable PDF files.  PDF files
may be constructed using only a `pdf.Document` object, without regard
to the underlying `pdf.File` object.  The `pdf.File` type is meant for
low-level manipulation of PDF files and does not necessarily produce a
file that can be read by a PDF reader without imposing additional
document structure.

Although the library is very much a work in progress, it is currently
quite suitable for producing PDF files, and it minimally supports
reading and revising PDF files.
