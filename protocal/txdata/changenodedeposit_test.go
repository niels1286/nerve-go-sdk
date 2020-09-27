// @Title
// @Description
// @Author  Niels  2020/9/27
package txdata

import (
	"encoding/hex"
	"fmt"
	"github.com/niels1286/nerve-go-sdk/utils/seria"
	"strings"
	"testing"
)

func TestChangeNodeDeposit_Parse(t *testing.T) {

	type args struct {
		reader *seria.ByteBufReader
	}
	tests := [100]struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{}
	str := "01000170ef34ddf98338698d5dac705b67c9193bb0227b00000000000000000000000000000000000000000000000000000000000000003ffe65b67d1bd2db183511c6abe920b3c0c8aa1285e162ece6da30e3f40d17b8,020001967160eed56d2fd4174094ee95e275d7cd686ec7400d030000000000000000000000000000000000000000000000000000000000203a4eebb5fc42a6210e105eb87e5e0ffa71272a841f5bec44cce557255692e5,05000159ed562a211c20aedb417642879d82c47e56a3a1801a060000000000000000000000000000000000000000000000000000000000b591b0d54e92d6ed404843b94b275283885dfd9cd22429a43efe3fadddfcb609,090001039271b80402e88c39e35f92ed282c934cab00bdc0270900000000000000000000000000000000000000000000000000000000006cd41fccb05452fd31fb94a0c96bd926c6999a9848ea0d5edbe6c7490af5473f,0100017f5e8e7365523294a1b95c28ee26e914d263f7ad00350c0000000000000000000000000000000000000000000000000000000000b6e41174b162841f19ce1266316172808b30cdd813bd362d12ec03afd8dd1db3,0200015b630b88702ac67c320165e08ff628996a9c81c640420f00000000000000000000000000000000000000000000000000000000002681f96c32ece8576adf73c35183a2508a9472f3b57271c8530fc3d07a0db100,050001bd3f20f614204ec3f5b0e2359bf04c20ea481b45804f120000000000000000000000000000000000000000000000000000000000819a8073b2ffa032d6fdbec963a96f87ee2a952b47b37d6eb2350763f03b481b,0900015ebac85b53af9488cb9d6efb2da624793c8db3c8c05c150000000000000000000000000000000000000000000000000000000000a8d1f1615bd7af75c6ebdaea55fe08ff32514b131cab37474a7bc6d430558493,010001ee9e2bd79d06a7144333a63b648b657975b1aa8b006a1800000000000000000000000000000000000000000000000000000000007f9928ab2848c890b155ce8ce073d0c0a4e0117ac9072059f6a6a7049a675bd1,020001c353212a7ad1c54e28d9238ac2b3d5374c067e8e40771b000000000000000000000000000000000000000000000000000000000052d391694411ecff5f44a212916d3c6d52a1d2cbe14d2856c4dad5e38f1f38a4,0500013460dce1047fa1f22c307d1b60003b405dddfb1480841e000000000000000000000000000000000000000000000000000000000003d6c4ce2c44385dc9bcf5079b134d7b0189e4d1a9c4e8b67dfac01e13eb5e90,090001ee9f2ed9b01d4eeec51426e7412894976e5e94eac091210000000000000000000000000000000000000000000000000000000000a4d037ad2a06bc106cfebdcab987a2cca7fbd165ba9408b6b99d17154d704402,010001e13030516be3eb8a595ace2b600e33e0e6b858c4009f2400000000000000000000000000000000000000000000000000000000006dbdad6233060fa9657126d52929a3c1b955bcfa2f944658c802705c6a8f9267,020001a4e5de9049eb2b9e85912e17f1151612132e823b40ac27000000000000000000000000000000000000000000000000000000000082a5ce59716680c5761c6b1b6935a306165e4b6b651af9d0a0659776c6da590b,0500012fdd6ea5a8efed82357360754c10f2c35e85d8da80b92a00000000000000000000000000000000000000000000000000000000003036cbd3d29af879e70d68b82b77a6516c66a878a928230482be1f3d032bc390,0900015c3a0c739b4b930cd180f12ba6adba610cb7ca7ac0c62d000000000000000000000000000000000000000000000000000000000053d77448d826ae8f3ecba746f889a5a4b0db78c6f71fe5d8297fd1ca19dd4e53,010001fcaa12910feda9188570a08420290c04ce8bb36d00d430000000000000000000000000000000000000000000000000000000000043989c9e03c6c1af15720d944a4752bd6dd3235c80b50529ba4422328f6d8023,020001ca1f8c7359f0bd6f3258afec721625b96b33c7f140e1330000000000000000000000000000000000000000000000000000000000e84864d78e33645a66b5627cc59413db6e2f08e4a1dd1222f4abf9d4a4281c1a,0500011b987ab934d50cc6659b8719c849299cd5f8726980ee3600000000000000000000000000000000000000000000000000000000004800ad25cba0675c01112b8751024de05d4366af3d091482bb89e18b086a30c3,0900013777a6f4e404b777b8d4803eb8bbabcaa03a045dc0fb390000000000000000000000000000000000000000000000000000000000cf273219cc846689cddba8466eea079050859d68583c09425ea5b1e12d3ee2b2,010001ae18a51b77e3b2bf79db0c611b71d64431f99e7000093d00000000000000000000000000000000000000000000000000000000007cd62b4f514fe7a37d1874b88ca0fafa59716d18af0908e8699096a18d8ff268,0200010b43b50e477e71022275203af2b858169d7d3dbc40164000000000000000000000000000000000000000000000000000000000003a7540ecccaca5fd5650ae040f757c2bdb97a30743f087f6f17d6ecd17bb0f93,050001143b10a5609d15e0257df3c457dfa85c0856d44d80234300000000000000000000000000000000000000000000000000000000000d0da5cd6480504731ebaaed54b767fe4f4ea232ce86c8f60fd9639385d70b98,09000111d3b3989ef138d29771900841960ec1f800ed0ac0304600000000000000000000000000000000000000000000000000000000009e4340b24e30f0c471b5219472a22ab2966951470659c8f8895496388b58ce28,01000117a04dfbb9136e968338649f2debed27613f62b3003e4900000000000000000000000000000000000000000000000000000000004ddc099eebfc4776fe8e5b75dd2546f9c013c2ef738eb2f2c033417e901c5ccd,0200014daee04ab48c797fd91e99a645495861e814b79f404b4c000000000000000000000000000000000000000000000000000000000078f94bf760d85fdb0a9aeea655a3401f4313ffa77f074ec38fe7b0683d9d23bc,050001a879499d2aba7a6e6a62b745049afa272fbab20f80584f0000000000000000000000000000000000000000000000000000000000f765bb141d42549e9bb40f41f5efa6e95acad7df3859f925f04fd842d79ceaa0,090001dbfb2e53fd7b3aa9554441623d360d7ba369fb81c065520000000000000000000000000000000000000000000000000000000000ec349930df726d637507a9f6f6c958532f04f187c5b7a89a7fed9641e07d1bff,010001fb55f71680f68ff3090d2edd62bf5ad23d7db93400735500000000000000000000000000000000000000000000000000000000000c23b3f777764f56b3bca589e08bef806c133af8707d21495dc1bf3f7b2c7cb1,020001e5a79f6abd98cb4dec1ec31fda7c40ed28bee3524080580000000000000000000000000000000000000000000000000000000000f722c1c28487913fc6b5319a44c4954c5544aafd3b7993696145b28e2436a318,050001f6d6509c047c64542ff38ee6ef81f9c02ce2d2d3808d5b0000000000000000000000000000000000000000000000000000000000a28743eaf5a74249aec65096dd488f97faff10186d4440d5acca006446833600,0900012ca18687e3dd8899e0ae58977a0c5206fa8abe95c09a5e0000000000000000000000000000000000000000000000000000000000d3a739cd1e1647755a3d6422eb79f2000a1d2c069710eb8d81cf044f14ceb008,0100010d711b39aac4754db166a0930b29f294c5de80e700a8610000000000000000000000000000000000000000000000000000000000c93a4d5405017288958f8c5db4a2faf1cd7bff6f0362010ceb50ad94fac12885,020001e6969d60b1a78f13594d69688aefa41ce6aff9b240b5640000000000000000000000000000000000000000000000000000000000f1531becd1745836b8e3dcf64a106071d809fad4f91349c18b339827c308c9f1,0500013ce069719def0587c4133db62d3bb2816cd24d3780c2670000000000000000000000000000000000000000000000000000000000ca58f1070e53dea2b7daa243327cd7241bbe9d0830b8d6326af077904162ed01,09000132ec82dcd91c7d09d56c9ddba184b1e8af5ea4afc0cf6a0000000000000000000000000000000000000000000000000000000000f3958c16565655267670720b150e1c520377445cd765348c72421cf474c8f013,010001b6abf3c6c67801516a100912143e7d6f981d0d6700dd6d00000000000000000000000000000000000000000000000000000000004b34f139b7c7af3e84fe21a9e1e351cc074e1454959cda53d54051cf2cfae74c,02000194a7033a346246d491ba8573d4a24f527b10239440ea700000000000000000000000000000000000000000000000000000000000f0e8b291331a0ddeabc5362d168e2a3d240845ce3cfd44310db09f5ee5f9c826,050001f5f49fe7120bde5c71afbeeb0a7604bce05e2b5580f7730000000000000000000000000000000000000000000000000000000000b9ea1f116a8c57fc364e9837a6ec0a1614b980262a6f960af9b26113b197a199,090001080f1628919d4aee1967c2f92c35aed6cf08b577c004770000000000000000000000000000000000000000000000000000000000189a98d900950d4b8f2cadb8afa804260c66c1086205d5deb5d6f9cca9d4078f,01000189bd68b85d3fe98faff8c2c85c246a149d56951500127a0000000000000000000000000000000000000000000000000000000000d0c8c768277d7f6f71e20070aacb0c9440155069fc323b222f0d88803c66ba54,02000149d95add02b9014c356113f61782dc177a62cfdd401f7d0000000000000000000000000000000000000000000000000000000000d7f384aac712db94caaa2bae7d91414a11d3572e7a91f51454b8daf2638648be,0500016a3b05cb44b1b41a8262f266458920ddf2b5757c802c8000000000000000000000000000000000000000000000000000000000005bcf7eb34743ff901a8981e3d078f461c55f91cf6a9e6a36a067cbc35a363cb5,0900014f56aeb0ed8222882d48c9faaf773c0aae81082ac03983000000000000000000000000000000000000000000000000000000000040f09bce90e24fddfae268395e0397ddcdb1d93ee126db05c47a3f7869fa9819,01000168719436dd380b4916a685f4767d229cd8961a30004786000000000000000000000000000000000000000000000000000000000008414bdd6ccfdeaea86c02773320b2be845a37d8aed3ca65f18a7d548fd71080,020001eec2bae36e051cfc326189ce5af9f9ab978794974054890000000000000000000000000000000000000000000000000000000000873dcc80c8a59dbca8cf3e17d1c6204c4b55915b10422639f814af769be6ec0b,050001e2dd2abb7c67eb4ba47fd15ce07055eecb06c24a80618c0000000000000000000000000000000000000000000000000000000000c097757a2f1831e827aa878571dd3d86b52c1c5e455f54fc9e1b2743f4565be7,090001370de470c3151bb0e05b575d16680b96417a9c79c06e8f000000000000000000000000000000000000000000000000000000000031b64e60b51ec7c7897ac8bc49fae92079afdf025cdee507c40d589c1fb1e27d,010001932c82e0b4779d3b081915f8c34d8678aa403ba1007c92000000000000000000000000000000000000000000000000000000000090884be4d83ad3b66c02331237d6aed11665edbbe36ad74d459b1b5fc232992b,0200017bef2a01291482bed47da43fa3a92bf68fc5d0ac4089950000000000000000000000000000000000000000000000000000000000d8478d3ddb1bbd711efe88d6a17cf7645e41d957e05c3e62dc8e36589bcf4daf,050001fcefa5d02785349199ef78446ac09a2f070347f08096980000000000000000000000000000000000000000000000000000000000ac0e4b9c31ce592ae16796797856b6afa13e678d7a1124106a06b26efc20f907,0900017cc06677646bebffbdfdce23ba0b3ca68b3ad2c0c0a39b00000000000000000000000000000000000000000000000000000000006d03cb6fba1776c2ca6b2430a249080814dba4699a2421c38e719691b0c074fc,010001eb967b35eda78d3e661f850030e9b3db24ae8c7900b19e0000000000000000000000000000000000000000000000000000000000939b5eb80c0b19e5f49e70b40bb5db18723d84ccf11903d421a23dbe8776ef4f,0200012c7ba28243a72b14469d58088b7e208175d25abc40bea1000000000000000000000000000000000000000000000000000000000054114eab177e087722eb1d381492529be99541ee1db94a3a12a9199759959b6e,0500016a3177b285d9942e26eb08d28aa905d78599975f80cba40000000000000000000000000000000000000000000000000000000000d1026424f4c6ccec14adc5048ac05027cd101843c1da08370a7d2158927fc16a,090001f198084c0037208cd06a537081082bcd6167233fc0d8a70000000000000000000000000000000000000000000000000000000000a91e47be4ddf12f44945b8449a169d5ae54fcb7c5c62e0669483965c02cbb5a4,010001b89d6fb7abcc3e1cb44e7a0c061a0a393203a7f700e6aa0000000000000000000000000000000000000000000000000000000000b0e8bb16796262369fe2c83d2ce9be67897668fb8adfc8c3f66835d4cce321e0,0200018f62ad809419538e01292a227146dd9954ab667d40f3ad00000000000000000000000000000000000000000000000000000000003e1aba1f28ef0dd7bc044262b12b4a50d247384bd9488e6e99368555b9d6ebeb,050001941985a7b51b6508d27a39d165a2e7f3dfa5ef7d8000b100000000000000000000000000000000000000000000000000000000008c6ee319ea8f4a75ca20e9a31e753c71efccb89808c6e6e7e1e676f58564d99e,0900010cecbe4b0c75e81080fc74a13ac0b4c640bab2f6c00db400000000000000000000000000000000000000000000000000000000002425f531c139786e852b3557392b422b8fd1ddfa8ce09c557169e8567327f8f1,0100019405899cd9de3f727a2ecdaff7758f46358b407a001bb70000000000000000000000000000000000000000000000000000000000b21b7bf23bd33d792a2de324d162cc9bebd62c67b6cc89a5bebeead4a552a3db,020001e33de60fc941513bedc595f52631d87dc6d9f6154028ba0000000000000000000000000000000000000000000000000000000000be84effcecd50f2649a5a9f3ddedf32d6b29786aee49f5233877cc0ccde8a87b,0500011dbae10dfe02db0ef2c431dcdbb91c99fb9241dc8035bd00000000000000000000000000000000000000000000000000000000009dbe2682a2c353ef0c9353f3c53a7ad7e01ac342c55c0120bf0130a571d67494,090001c9084a09492d2c27d616ca9f3b8f4216bebc50bcc042c000000000000000000000000000000000000000000000000000000000003ec097b14c5b60a956e661113c700563c0a537716fc3854fac8be64f79490ca4,0100013cab9d4e6e14fdeda5e41763eeac337200fdfb8a0050c3000000000000000000000000000000000000000000000000000000000028d0677a2721aea8b3cd2a837cf0fc4dc6dfd15954fcd7e23ef52d64c64b32e8,02000133ef91e6a3e169bb7e00c1af0669a79c6949a2a2405dc60000000000000000000000000000000000000000000000000000000000a03b9259b045a4ec22590ca01624e4b9d21db9f76b478cb9e645c9b5649839bb,05000179ce41318201275ebf1fcc1c7d6b6bdf439845a1806ac90000000000000000000000000000000000000000000000000000000000d15e2434bfde28e1ff7516a8a106880c2a2e010b9248108a1f70d1d78be40a8a,090001a05ae372d2dccfd79965727705dbaf77b6f71d49c077cc000000000000000000000000000000000000000000000000000000000054b4dcc3a277666ef3186c159f213f87b912ad6154d7fdc4c1c1257a6e5b747b,0100017967d7b5dd68cd56c8a0fb1ebff717e986e931630085cf0000000000000000000000000000000000000000000000000000000000aca4867d3a3564bf7ec649101c8cafe688d5c943d557b0cc4329b27e8727cef8,02000136b6140a82e40d54ca2b75cbaa955f9d3cbdbe914092d20000000000000000000000000000000000000000000000000000000000b2e142836c4ebc1267c65da4638d9117869d1bb8576d67db1a5dc98cba439a4a,05000155fa2993ca7ab908301c9a71fbfa69fe473b7d6d809fd50000000000000000000000000000000000000000000000000000000000721fbf8db1a3d3eeb6d438b0faa2c43b6a760411386043bf0e816ca09eee9e70,0900017192b58663d5b5afb7e73f1833bc43643ca17150c0acd80000000000000000000000000000000000000000000000000000000000d807aa9ba42f2b30b1bd325298224d00a583b79699c27831c2cd39a3c7f249d3,0100010e86a57b2d9322737ecb9664a1a940841cfc5d2800badb00000000000000000000000000000000000000000000000000000000004af3bd609348bb0d8d33687a48a46d5fea2ca3b646b8a5f42c3da66c6aaf5557,0200017eaf1d141f631ae271ff319f86ec6e07aa1a221040c7de000000000000000000000000000000000000000000000000000000000090165a1bba95aeeb9f9c2da640e7cc6e01ec533cb396d2c7356b1306c8c72772,0500011fdde332ff7166ad2fbe2414d888384611ec161e80d4e10000000000000000000000000000000000000000000000000000000000226c5a5c55c78b39e655e143a30e50ddc535827d544b8f7b7becf0776b677ff5,0900019203c0f7e20ccd55415737736a3114ac84f972b4c0e1e40000000000000000000000000000000000000000000000000000000000094b398ad2c6ec43a3ea0486ecc40865ac60d7c15bf6c6a43305ae59c87e0d66,010001de87238478d50f768ed3e3b1c2c13ed52277205400efe700000000000000000000000000000000000000000000000000000000003781554a832d16a4eb5897294cb5b801f1d4f28f9e24dabc5817f1a62ea31817,0200019b0b58a1f43f36255b297c62f88429d028dc23af40fcea00000000000000000000000000000000000000000000000000000000006d390fc95222de16ed469f4e354d0137d9f6cc706ea8d8e83561c821df6e980b,050001ccc9ae44576601c8ecefd96564aa3be618bba2028009ee0000000000000000000000000000000000000000000000000000000000f46eaa9783642e9711f95c9c84c6fff8ee87343a471f60e27e52638025703988,090001272b6c1f66ca80eaa879ff168fbe067bd58a6d8dc016f10000000000000000000000000000000000000000000000000000000000f1a3ee56d9f0d773b0405ffd6e5e5269b450edbee7c871439727bc861c6f6d92,0100018e610e95bc9cf2e6318998be42a4f1513bec10720024f40000000000000000000000000000000000000000000000000000000000cd80d0fc1091faeb256087604baddde6a1e37abc18703f4926d0fe88fae33a7a,02000196cffddea6e540cfaaea3259377b72aaf24ead584031f70000000000000000000000000000000000000000000000000000000000df521588fe2c777289ac816eee7498382cd62ac7ec3a20828888ffb2696a14ae,050001fb64ea9014177160352d477df7428185f52e1308803efa0000000000000000000000000000000000000000000000000000000000256ae9e1777bdaf51aed7a237b8a8b256cd6bfb40dcecfd06c43f41e874cd261,0900013c47e300e9f575a20216db4b65a00460ee1f8383c04bfd00000000000000000000000000000000000000000000000000000000005221437e57150612245ce8869837af92c1295cb54db21713da48aaced7896df5,01000194c6626031673d990f5bfa9c076d1e0865e4933c0059000100000000000000000000000000000000000000000000000000000000205c1327bc8e38b159755671e69332050130ab031043b033aa85a0abb1d1d436,0200018e193c692f0859f6b20caf38e69ab46f8b2f853f4066030100000000000000000000000000000000000000000000000000000000647404dc61138c2ccb8f26e340c4361e4edc4a2041e4a0b7f58dc8358b277677,050001fdc971ffd774fdfcb6b9ffc5f1901c2ca247fd9d8073060100000000000000000000000000000000000000000000000000000000604e4f9727c1498f6b4e73e9d6aee82a6e2b1e86ddc33dff679827398a36f049,090001b945236033ed9ab4c40d16ce0aaf60061cd90871c08009010000000000000000000000000000000000000000000000000000000027aa11097e6a43ab149b7c57137d059c74c12946233b5c1153d2b10c73beaf40,010001c8a2f42bb87066d3e2e6bb899b837962208ca829008e0c0100000000000000000000000000000000000000000000000000000000a85842645dcf63d3f05fbdfa049dc4048fbcd497ed8447f7a63bbe5fb5e84642,0200012d84ef3670588d37ecd643910009f6c1f7a7cd01409b0f0100000000000000000000000000000000000000000000000000000000d49816e3c3c4e9c538a6ff51dde6e7ad8acfa245567f656df5d8741b855fc335,050001359209e4dfe5a5c695c9effaaecc64aaa398cf1280a8120100000000000000000000000000000000000000000000000000000000e4c3a17dd6e19828d6e614e1f7de873e12cf5771f4d15be302c66116eb3a1a6f,0900019fa8768f4dfc656c2b1e11b15b27aee1c53cdf58c0b5150100000000000000000000000000000000000000000000000000000000b765ae648c0b6a6f0a7d724a836c1e9314d77d2101b3f749e6c87f1dd99babb5,010001ab1e9ea01793f33983ef34365872450523f0992900c31801000000000000000000000000000000000000000000000000000000002ac37ced8505a92b24e4bb2b0565659e57ccede5e5359aaa63906c5be0654d2b,0200013a5aa7b3200c17685ea61978320d07384cc9db1440d01b010000000000000000000000000000000000000000000000000000000051b8ffe2372c2b8f02d0ac1a5fa499f4bc230d762cbf16b183176f472b9544f6,050001a9b609ed5c535d2debaa66261f3bc23a55db7e8180dd1e01000000000000000000000000000000000000000000000000000000002e61b757447a519701bbfa247a2f24bce23163bbc9657ee50b494ecf366b7e5f,090001a172dfcb09fc7e4bcaedd8c388f480574ca87ba7c0ea210100000000000000000000000000000000000000000000000000000000c96734b8aa44a4ca6fe881d4e4807b70a779b4b8ffea57eb09a3f0c35213a47d,0100010847a311b810aa4873763dadc332ff46f14aff9400f8240100000000000000000000000000000000000000000000000000000000871a4a1ce4988d5e6eb70412b3e2b51a1bdb81a58d6027630d43b7472fd01d63,0200011a13e83ee945f7c33fc53b003ef4134ea229b49b40052801000000000000000000000000000000000000000000000000000000003dc4ceebcf5e3c23cf22f2046a68cdddf925f10a5c4c977a393eb0a646907475,050001a4be7b43192117f6d85dc978fba5c20235f9315080122b0100000000000000000000000000000000000000000000000000000000a395d6ef43ca8f04fddc8a7156e894b5f854977a4623aaec9ecba1713fd41a37,0900016490144d338c501391dac90f558023449eda1022c01f2e0100000000000000000000000000000000000000000000000000000000965b07f988c191598e8f41c10f884a20c4689fb62fd6015757cff5ce4df0c885"
	hexArr := strings.Split(str, ",")
	for i := 0; i < 100; i++ {
		bytes, _ := hex.DecodeString(hexArr[i])
		tests[i] = struct {
			name    string
			args    args
			want    string
			wantErr bool
		}{"test" + fmt.Sprintf("%d", i), args{reader: seria.NewByteBufReader(bytes, 0)}, hexArr[i], false}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ChangeNodeDeposit{}

			if err := a.Parse(tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			bytes, err := a.Serialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := hex.EncodeToString(bytes)
			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}