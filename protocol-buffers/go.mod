module protocole-buffers

go 1.22.0

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto