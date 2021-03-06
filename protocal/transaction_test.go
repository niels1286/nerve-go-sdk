/*
 *  MIT License
 *  Copyright (c) 2019-2020 niels.wang
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 *
 */

// @Title
// @Description
// @Author  Niels  2020/3/26
package txprotocal

import (
	"encoding/hex"
	"fmt"
	"github.com/niels1286/nerve-go-sdk/utils/seria"
	"testing"
)

func TestParseTransaction(t *testing.T) {
	type args struct {
		bytes  []byte
		cursor int
	}
	txBytes, _ := hex.DecodeString("10008850055e00fd5d02010001c47136edbccf65ed4e87e72ee8fc108fdde1cda9010002537f1ecbcecc1f0eca4b17260134a07c600bf7ef000000000000000000000000000000000000000000000000000000000000000080b506000000000019000000000000000d616c6c6f744561726e696e6773000309254e554c536436486755624d637a4238635051714e664551613177316e545538546143415241254e554c536436486765436e796b47534136694536396e4853654d615032505a555071315559254e554c536436486763634b5a4c524b46697541343864716235704534703537683177757037254e554c536436486767527a504c324e663832695574653837417942484236717936514d4377254e554c53643648676462386f367943465956527659745068514d7645706276564155774c38254e554c536436486762467657515568794b3679394e79555370643657413438366d4e6f7233254e554c53643648676a5135666f4b5276645973564a75734c61354b5963636d656245517973254e554c536436486765516152616e426e7a6e736454333552596a70537a4345464443566967254e554c536436486766623772564e644c3453744e42587533486f69654d674331316d6a774809093138363236333430330834363439333732350934363632333731353509333631353531363336093130343830333634360a3332373135373439333609363632393331363531083535363033323138093436353637373839380906313134313230053238343838063238353637320632323135323005363432303007323030343634380634303632303005333430353606323835333336480117010001c47136edbccf65ed4e87e72ee8fc108fdde1cda9010001002040a9000000000000000000000000000000000000000000000000000000000008fa878eb746e5bd3300006a2103728fbdd68ad3259e5f50d9ae0d5e553a01e8cbaf5b58703082e79bbea3a41e1a473045022100ee079a7e7c630e18dde1205fd350181505fedf0bd354a06eef25526f33302b390220793f2fdcbacdf618cecfeeaf90fcb96c00f5fcc45ff94f11b93d599282f71965")

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test parse tx.a", args: args{
			bytes:  txBytes,
			cursor: 0,
		}, want: "558cdf35fccc25c4e7a01b5891c2d942bcfe4bc9ae981dd9ac38b429eb39c505"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseTransaction(tt.args.bytes, tt.args.cursor); got.GetHash().String() != tt.want {
				t.Errorf("ParseTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_Parse(t1 *testing.T) {
	type args struct {
		reader *seria.ByteBufReader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{reader: seria.NewByteBufReader([]byte{1, 2, 3, 4, 5}, 0)}, true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				hash:     nil,
				TxType:   0,
				Time:     0,
				Remark:   nil,
				Extend:   nil,
				CoinData: nil,
				SignData: nil,
			}
			err := t.Parse(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				fmt.Println(err.Error())
			}
		})
	}
}
