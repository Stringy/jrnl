# journal

## Note: this is not intended to be entirely secure, it is intended to provide some level of security against casual snooping.

journal is a tool for writing and manipulating a diary/journal from the command line. It stores a single entry for each day in an encrypted and compressed file wherever you want in your filesystem. 

The tool currently has the following commands:

### init

journal init /path/to/new/file

initialises a journal file at the path specified. It will prompt for a password to encrypt the file with. It is encouraged that this password is long and secure, to ensure the security of the file. 

### add

journal add /path/to/journal/file

adds a new entry to the journal file. It will ask for the password of the file and will quit if authentication fails. Otherwise it will load the journal, and check if there exists an entry for the current day. If an entry exists it will allow you to edit that entry, otherwise it will create a new one. To edit/create the entry it will check the EDITOR environment variable and use that. If $EDITOR doesn't exist, it will default to using vi.  

### list

journal list /path/to/journal/file

Prompts for a password, loads the journal and simply lists all the entries in the file.


## Plans for future commands

- Serve. The intention is for this to unpack a journal file and then serve the content locally to be viewable as html rendered data. This will allow for far simpler browsing of journal entries.

- View. This will allow you to view a single entry of the file. 
