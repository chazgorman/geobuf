package geobuf_new

import (
	"testing"
	"github.com/paulmach/go.geojson"
	"github.com/golang/protobuf/proto"

)


var vals = [][]float64{{-80.214562, 39.722209}, {-80.214657, 39.722396}, {-80.214843, 39.723198}, {-80.214837, 39.724068}, {-80.214739, 39.724274}, {-80.214631, 39.725024}, {-80.21468, 39.725342}, {-80.21488, 39.725888}, {-80.215006, 39.726014}, {-80.215262, 39.726269}, {-80.215999, 39.726711}, {-80.216555, 39.727157}, {-80.21686, 39.727487}, {-80.217047, 39.727899}, {-80.217151, 39.72824}, {-80.217104, 39.72881}, {-80.216884, 39.729378}, {-80.21635, 39.730418}, {-80.216273, 39.730568}, {-80.216228, 39.730975}, {-80.216338, 39.731524}, {-80.216544, 39.732006}, {-80.21691, 39.732364}, {-80.217309, 39.732628}, {-80.219144, 39.733516}, {-80.219498, 39.733732}, {-80.220198, 39.734404}, {-80.220879, 39.735298}, {-80.221058, 39.735533}, {-80.221408, 39.736025}, {-80.221463, 39.736107}, {-80.221734, 39.736573}, {-80.221738, 39.736591}, {-80.221864, 39.73707}, {-80.221867, 39.737267}, {-80.221892, 39.738659}, {-80.221936, 39.739471}, {-80.222007, 39.739985}, {-80.222066, 39.740175}, {-80.223055, 39.74256}, {-80.223267, 39.743234}, {-80.223395, 39.743875}, {-80.223437, 39.744605}, {-80.223422, 39.744884}, {-80.223416, 39.744998}, {-80.223346, 39.746255}, {-80.223172, 39.746815}, {-80.223023, 39.747097}, {-80.222815, 39.747356}, {-80.222444, 39.747666}, {-80.221809, 39.748196}, {-80.221703, 39.748276}, {-80.221598, 39.748357}, {-80.220978, 39.748866}, {-80.220364, 39.74937}, {-80.219497, 39.750167}, {-80.219348, 39.750323}, {-80.219098, 39.750644}, {-80.218902, 39.750979}, {-80.218747, 39.751573}, {-80.218751, 39.752086}, {-80.218796, 39.752328}, {-80.218818, 39.752444}, {-80.21901, 39.75287}, {-80.219317, 39.753294}, {-80.219808, 39.753728}, {-80.219926, 39.753821}, {-80.220321, 39.75411}, {-80.221006, 39.754664}, {-80.221156, 39.754786}, {-80.221495, 39.755132}, {-80.222028, 39.755874}, {-80.222291, 39.756361}, {-80.222725, 39.757647}, {-80.223267, 39.759193}, {-80.223352, 39.759354}, {-80.223394, 39.759489}, {-80.223511, 39.759859}, {-80.22355, 39.76015}, {-80.22359, 39.76044}, {-80.22352, 39.761136}, {-80.223342, 39.761705}, {-80.222967, 39.762512}, {-80.222778, 39.762953}, {-80.222481, 39.764355}, {-80.222485, 39.764733}, {-80.22259, 39.765106}, {-80.2227, 39.765312}, {-80.223593, 39.766273}, {-80.224527, 39.767457}, {-80.22491, 39.768156}, {-80.225213, 39.768901}, {-80.225204, 39.769409}, {-80.225137, 39.769654}, {-80.224937, 39.769984}, {-80.223607, 39.771721}, {-80.223191, 39.772308}, {-80.222733, 39.772971}, {-80.221743, 39.774451}, {-80.220086, 39.776907}, {-80.219642, 39.777601}, {-80.21911, 39.778224}, {-80.218956, 39.778391}, {-80.218015, 39.779173}, {-80.217186, 39.779865}, {-80.216548, 39.780397}, {-80.216464, 39.780467}, {-80.216136, 39.780767}, {-80.215832, 39.780987}, {-80.21567, 39.781123}, {-80.214983, 39.781791}, {-80.214563, 39.782303}, {-80.214141, 39.783066}, {-80.21387, 39.78364}, {-80.213363, 39.784822}, {-80.212761, 39.786174}, {-80.212219, 39.787266}, {-80.211468, 39.788353}, {-80.211335, 39.788539}, {-80.210167, 39.790171}, {-80.209946, 39.790537}, {-80.20952, 39.791243}, {-80.20895, 39.792516}, {-80.208048, 39.7945}, {-80.207743, 39.795277}, {-80.207555, 39.796097}, {-80.207568, 39.796672}, {-80.207593, 39.796852}, {-80.207605, 39.796942}, {-80.208259, 39.800099}, {-80.208261, 39.800735}, {-80.208184, 39.801196}, {-80.20723, 39.804022}, {-80.207086, 39.804714}, {-80.207066, 39.805311}, {-80.207247, 39.805925}, {-80.20739, 39.806172}, {-80.20806, 39.807122}, {-80.20918, 39.808585}, {-80.209402, 39.808925}, {-80.209624, 39.809264}, {-80.21012, 39.809882}, {-80.210301, 39.810107}, {-80.21041, 39.810243}, {-80.210582, 39.810484}, {-80.210659, 39.810781}, {-80.210772, 39.811075}, {-80.21099, 39.811305}, {-80.211285, 39.811414}, {-80.211836, 39.811511}, {-80.212159, 39.811656}, {-80.212371, 39.811907}, {-80.212452, 39.812275}, {-80.212571, 39.813316}, {-80.212653, 39.813726}, {-80.212736, 39.813959}, {-80.21323, 39.814736}, {-80.213304, 39.814896}, {-80.213388, 39.815234}, {-80.213302, 39.815783}, {-80.213376, 39.816146}, {-80.213664, 39.817035}, {-80.214152, 39.817777}, {-80.21437, 39.818028}, {-80.214425, 39.818158}, {-80.214526, 39.818396}, {-80.214489, 39.81863}, {-80.214467, 39.818772}, {-80.214176, 39.819936}, {-80.214149, 39.820328}, {-80.214373, 39.821069}, {-80.214355, 39.821335}, {-80.214045, 39.822109}, {-80.213723, 39.82295}, {-80.213392, 39.824185}, {-80.213124, 39.8247}, {-80.212926, 39.824967}, {-80.212438, 39.825468}, {-80.212266, 39.825644}, {-80.212221, 39.825687}, {-80.212153, 39.825751}, {-80.211353, 39.826509}, {-80.210875, 39.826809}, {-80.210464, 39.82702}, {-80.208427, 39.827905}, {-80.207891, 39.828147}, {-80.20699, 39.828553}, {-80.206563, 39.828962}, {-80.206328, 39.829531}, {-80.206108, 39.830249}, {-80.20606, 39.830872}, {-80.206132, 39.831358}, {-80.206145, 39.831446}, {-80.206195, 39.832956}, {-80.206223, 39.833507}, {-80.206239, 39.833807}, {-80.206377, 39.834253}, {-80.206412, 39.834364}, {-80.206458, 39.836312}, {-80.206469, 39.837273}, {-80.206386, 39.837619}, {-80.206259, 39.837873}, {-80.205936, 39.838365}, {-80.205392, 39.839167}, {-80.203674, 39.84132}, {-80.202788, 39.842544}, {-80.20262, 39.842872}, {-80.202026, 39.844199}, {-80.201819, 39.844817}, {-80.201693, 39.845137}, {-80.201687, 39.845625}, {-80.201742, 39.846001}, {-80.201857, 39.846484}, {-80.201789, 39.847015}, {-80.201525, 39.848364}, {-80.200512, 39.85078}, {-80.200453, 39.851096}, {-80.200384, 39.852495}, {-80.200142, 39.85414}, {-80.200134, 39.854197}, {-80.199986, 39.855744}, {-80.199926, 39.856372}, {-80.19995, 39.857389}, {-80.199847, 39.857831}, {-80.199707, 39.858073}, {-80.19942, 39.858914}, {-80.199467, 39.859455}, {-80.199614, 39.859791}, {-80.200183, 39.860703}, {-80.200377, 39.861283}, {-80.200469, 39.861713}, {-80.200875, 39.863604}, {-80.200747, 39.864206}, {-80.200357, 39.865205}, {-80.200308, 39.865331}, {-80.19991, 39.866727}, {-80.199745, 39.867251}, {-80.199559, 39.867559}, {-80.198822, 39.869298}, {-80.198071, 39.870487}, {-80.198, 39.870738}, {-80.197886, 39.871508}, {-80.197701, 39.873858}, {-80.197669, 39.874267}, {-80.197665, 39.874316}, {-80.19719, 39.875096}, {-80.196755, 39.875549}, {-80.196555, 39.875861}, {-80.196484, 39.876123}, {-80.196331, 39.876518}, {-80.196311, 39.876569}, {-80.196108, 39.876929}, {-80.195605, 39.877616}, {-80.195315, 39.878558}, {-80.195232, 39.87881}, {-80.194979, 39.87958}, {-80.194681, 39.880488}, {-80.194504, 39.880781}, {-80.193921, 39.881422}, {-80.192096, 39.88346}, {-80.191504, 39.883879}, {-80.191452, 39.883916}, {-80.190092, 39.884847}, {-80.189823, 39.885116}, {-80.189596, 39.8854}, {-80.189083, 39.886042}, {-80.18762, 39.887268}, {-80.187146, 39.887711}, {-80.18698, 39.888075}, {-80.186937, 39.88817}, {-80.186694, 39.889092}, {-80.186444, 39.889902}, {-80.186344, 39.890446}, {-80.186299, 39.890692}}
var feat = &geojson.Feature{Geometry:geojson.NewLineStringGeometry(vals),Properties:map[string]interface{}{"shit":199}}


func Benchmark_Make_Feature(b *testing.B) {
    b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		Make_Feature(feat)
	}
}

func Benchmark_Write_Feature_Old(b *testing.B) {
    b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		feat.MarshalJSON()
	}
}

func Benchmark_Write_Feature_New(b *testing.B) {
    b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		feat := Make_Feature(feat)
		proto.Marshal(feat)
	}
}



