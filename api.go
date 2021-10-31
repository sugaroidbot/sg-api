package sg_api

func Listen(w *WsConn, cb func(resp string)) error {
	defer w.conn.Close()
	for {
		var data []byte
		_, err := w.conn.Read(data)
		if err != nil {
			return err
		} else {
			cb(string(data))
		}
	}
}

func Send(w *WsConn, message string) error {
	_, err := w.conn.Write([]byte(message))
	if err != nil {
		w.conn.Close()
		return err
	}
	return nil
}