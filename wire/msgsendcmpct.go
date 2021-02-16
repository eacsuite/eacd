// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"io"
)

// MsgSendCmpct defines a bitcoin verack message which is used for a peer to
// acknowledge a version message (MsgVersion) after it has used the information
// to negotiate parameters.  It implements the Message interface.
//
// This message has no payload.
type MsgSendCmpct struct{}

// BtcDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgSendCmpct) BtcDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	return nil
}

// BtcEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgSendCmpct) BtcEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgSendCmpct) Command() string {
	return CmdSendCmpct
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgSendCmpct) MaxPayloadLength(pver uint32) uint32 {
	return 0
}

// NewMsgSendCmpct returns a new bitcoin verack message that conforms to the
// Message interface.
func NewMsgSendCmpct() *MsgSendCmpct {
	return &MsgSendCmpct{}
}