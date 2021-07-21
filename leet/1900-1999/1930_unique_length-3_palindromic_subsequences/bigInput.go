package _1930_unique_length_3_palindromic_subsequences

const (
	tleInput1 = "gywghdivwpgymbjqfbzelzsnmcezvvwimzvcdujeoexfaqmsphzsqwxtytjikvcddfzrmgqziiwqyfayezdzlkkffstwsxvqltxbiyplvvzypskmqdqsuscxvwgbbwargaveyiikkvscfvjmqduufoepeuegowaukvsjmkiozyoclterhrggszwyksguoyodfsrquhurhhjruuthiseqmqpgjbbfphxdkhkydtjgtyxxsuvlvakmqrbqrogpsgoesiyfdptupeynrvnswsagnnyijkdccklnkzkbyskemsivvyhwkvjswplfahkfoihucrdxzuqnofdbjpboqgzsqeseyibmkulvxrwfoojkhraayrmiwpkjvjaccjxbefrfskpmlqxhaqgedqguytmawjzekztquxvwjakyhjjnfkmdohpfydsnztswnnckryxsnhifirtlywtluzqaqmgdawrvwbqhuxlpfxmovbgjcjptwevolhbchzgrwxnmojuiypvdcqltabohaswrfuryiuphxrveyndwzmvhlglxttlikrchkqscknpjwixdnhoiqhtecdceyynsuouhbcduqvmscimlzwhdubteisewmaztbenultkhieimykwqvrklqxaueldtgdyqznkrfzthfljxoorrysispbdhaylfrroxxqhrrhzpglcgmpkqvlxuwgnaraqpnkmudgiiadeedokauinfzexfmfggxwwnoyrgoqhyerfydnslwkpocncnfyxgfyssreuijhleabgnprqmggxrpwffgwjbfzvwdbxzbrvugqlglrmjnplmcttmggznmzzenytxhriphvjcpmtndrbpkhaopegdllwsxlbmsvaciislwmjbkswgppnzumdkqnvsbaehtwfraxofdnudiirzfqbvnwathcxlfdsjsuhxqztcqmjwofsqdvnqdnwsoofpqkayvoxxtowpxtommjpnkwgvpdjbirlzixftztndnefdqmzumxjciymyzxdbdghsjwokljpdlcoskjshksqlxtpugrnlqncstkedcmtkagvaxobuamsozdvvekmgpmtfeitcfqdmwqwgsshhmenmpehlrrxbikkkopgwrjflzwejvdzpxyxonuopezecekixsojmaykeivcfidznhfruexgscexdbgbzonqdmqrlumffcetbpwmcmxrygctcpjowzovbdnzztjuixgjkhrqxwsbwbmwoszqewxmssbxigsdetyeghuvyqgaooiumlgjrnoyicfalkmxasvylmdqsmhakhrjgrecyjirinobqovnvsadvgjbruwzqeqytjgzmrugismpceybzovifqtgrzceinspjivfmpmoncsnbkcoyxohhmucqzthlxymknqiisriuwhygemjcysijmbcqvzektocbjlvdbikvkthskkrqofemivfnsfrggewnyrnyqtyuligpxtaadncbfekkconnkgsxpbzdmmkawphtjtevojffipxzbpwtxjhpufavbtwbltxiafzxgdtaowfyyfwwfpjcmbatqlcchtkjroxesrwydnavhqvzfjxrzbbpuvgrsjixijycojomizwzjrogdyaatnralgtlwybycwakmmdbbhlhjysynveqyquhccjwmpdfwqtvmajrqdgtzzteoxrrimjdhmbnsrldwwjfevhoevtedbwfnlxrdywbxkljuanuztxciefmvvdyhhfjbovpuvbpkasywkhnjnheemkjisvxidemapdjsvihwtelbyjeuzplreaxfiwlfebpnsfdfbmqftfhmmlbuyinnluzhvuwxpbrwofagmnincgusqqdowfnkqlakcsyjubyuqdnupjcjnamtldlacekfdtewekjwbcqzkagidkodybvhlpadgmcysudyjvdevufhptqisyfzcdytqhzfvunkxlabnnwjrhhqcczdhlypmgcjwaislwgxcorrqhkrehiygqoamdavtwsjzhnbdlmjaczllmilesvghxdualhdtatoafhfscrnelsihvlgfgwdinnzvqtxexwwfayeygkpfnmcyxiwwssdecrqesbtadrhzzeroqojtfvudisgfjoqewnsxqpmsisvsxhiispmewpwefnorcykghuvpcgcnhultrgobrrwylqzkcezmqqhpajfcngjeofrivcsdsxvpzmaolzvrleoofuhcjuvptfpjepuqfmfiqebyctsrnlgvulzvxfabccqgcodtmycsmbcforgvjhppwphxpqrddmecdxihceerwwnqxqactirfcodaqpekvdwdqgrouqxkwcinallbyutjrpzirvkdqewonvnmyrnuhtnybedhjjxfmsddiksujcvzunbicveqhtxnztblkgohgunxjsdunotlunithlifffkbpodtgoytzqswuhdzyekftafcblnolyurbviiikrtffsrcxdrdewxzebmtzoeaqlytmbtefuvoyfmmtpcpprdeagltimgaoiyyspczzpofzptqodlfgaaryovxttvfsbsbmkpxxkxsnnqasphxuxpegtmnswsnpqdbuzourxmfsdxwxmzlgaqtgayfthdghgtcmigtbfdrebblbwaqtlpvygicrwzhjirfbjevhxymuuvxpcmlnsoksaepjsqbczibasdxqfkcmovidxxvnzixyvhojbhqtrctagutmvxvovidfetwhgriqwrbftadvidiwaeicvumbnwvfnkaccfvfspctsovvrlhkhgzbvfwrvsocdrzgawddynatuhbowggygbvvwsjwjkowbedqaestyrosljfeuhiorwpgcecgspnfallbrkbtoufcyuipofcjqgqmyfmivyocergaxraifvioniqeoygznmzmsswxjqyavuichdnnbsyswmqymqpvtfrvajczjqagjstghghcxondkmhgmycgydazbcrapbqaizmuewxghthyihgtxzxanxonnckglcxrnumuvxdjjhnlnktqxvlbyenttmigwdzhejoqgobgsdnnqssvtugrnefqvxbormcwxuybfwhadbhkgjlugbwijvreqdgwqmlcotxodrywqljnjcgpnvezhrhkwptdjmwzmrfhujymehgmqbznuguxykdbejizhghfkfycagbxnlhafvpgyrrhupffswwecidudkvyovxytbfbckfudtluljhbgmetixzlupypsdkuhorbkbzpyesqvowmmfqwfecyottdqwnliryjrvqfhxnuaqlrlrwegquyqkketzyoowoowgultllqkwdgauljgqhsygnbbidodvxmamwxrsnkxekblbeqdxaezrhkcsimvokzwjmbxonfofzsgqsekpltojlwefiogbrxlsyhetdgwrupksvxcpyskbgssbyxomjrbeaxrlufegidpczzrkqctxegvvpcwokgescuisokwyapwknqmrmbssscswkhvnyyphbrvapsoodumbhbccrgopwknnnoapduumuizombssqftyzpqboxtufyaaxgyfqxssqpihuvtyvyazwokoexefonqhecxlhmwmeisgiklechdslygmskchfqfatxnhwefboxvgfujsawrcuibzytdpfpdtfmgisfzkaehadluidjdlrdtijqgjwqenfshyvujkxcnbzkvjiuwcspicwivmoilcgcyddnvbhworqgqivhjuigninvyqsbohponoppqvwhhabupyojhiqcooaxummbyncfepfyumbdmdghsdvlqudfgsnouuurhcehddvcvviqixtartsarknwmgmrvmaqqysquaembbqyuahaelilsqwxtfbivcpgceznllihxzvufgtdoncnjcppprmzatjymzusmvxdumuczuhiavtguqqvlqhslvfdmgwseecfhaknrqkjpxxgjyskjnjxxtchftrbtobikclmbtbdcdriebkghaslhxxzaxufvqmnmmjsumsxeuegmireomkxudtdzevpwolovrvnpsasdwjzshalveinkynawhbyjkfwsutxbeicunsnwhtakpzxxdgjmajgofmpyryiwjyeziooksjpjhwibhicuakjkjnpoqfqqczqgejmecokzveorsnwgnkjlrrcsoyyxzroejcuiofjtwhphsfgqlyouvulzbtykqxxpjrjnwkttbyeyjcwtmahoalmsdzmavxlupkjudmwqipsnfoovvraknchmozyvsufwillyzqnileazqwuwkpcsuhpfupkmsjxhwcucyjfblzoqlqicmkfuvbbdpxvwlkyyeuunvzfvexzfouefpytxbmrlkdjkchsuaevoyttgjcuharjaqrjeifaimqwnuqdwwzakusqmlqhdbnyfhdszymxbmhskzqarsrdmkdymdbkifwupotrskkorvkketdzvowrdgasaojgomfhosktzoglatowyihczvjwczkcezpejcrwmjbfukwnaodrrwibcfnfrphqcrmewpghuayjbjmctlkwmehekiagwibyouppsegewgwatlxciqsdwmkrvkotojwrdfuusejtucenakvlnnxffnhwdgxxrjbfjtpiasumropsstyptvbxvsdczmelbmkremdscukcnafetzphfdpnvgbivzrsnebvcpdomellfythwnrnapayvtqhswtpmudgzpzenwyfwlqgrisujhefgbjonjibnpqccanlngivuwphvsoldvgwziytotunjdmcvcwrdlcpougzktfksuwvirkqsabhhdrhoetjybuqcquywhzewvxdfjxqhzxranjiroehfryszxzcebskpxfghojdqourmvmebfomdnrdvtslkyjfpznafglphjgvakyphieepaqobqvmceoeaesvzqixefldtbvbcwytytzlhbkioujxnihmmjgduhlhkezikyfykvxaveiwscqrzcwenbasbecaxngzqkbidzfpltslynsfukwhnpgvzwffafaljkwgzegqbddnqchtcdibrnnthhjvqlmqtgrfyscgywjgtdrwhtkaphxuayxqjpjcccujvhbyzuiagohhwxtxvyjlwbgaztieahqjmeshadcdepfeayhwcprmsmhrglfslicftnmpnwjtwjyiigiqqiichkbgmxgznokawqyhtjvrorpgmsobijaowfihhaisqzpimwxmsxpjwoxbcnhzejgmbxexigfojfawdghobqyqkkehpakwxfpkwjilesbjqezufpqubebbdrzfmeszoxytuounwroyxskklrhffcipfxbedhtcxbhzfjqpgrnhpvwejncxurjhkwaefecsdrywbcfjrjbcovymrjwhkkhfmtagtcabqewxqfipgynjkddifltyaulzkkxmyehobbtvaagauuljayzdrjtohcdnnemtzmuzdzqbthwooyclcbkztrntoqqjgnvlpcjxipgtfbxmfxehepjebpqghhvwuhgglidcayjvszqcwpynelzokarzggjbrieigrhlfmocvvzqvhbmbccefjysrztwdytgqaksgnklcvtbkzxjqyudogbfbcwhmvwerpzrhfckuchhvoaiwaifturpuazmqsspwqpultsfcwawyltixsydwbwlboxlvoitoazzvdnmhbwixycjanpicocjuwznfnbudgxwllcyohdcuftzropujqmzvqpjacohbdtwyenlypybnxskvhowwtkxyduozjlewsojelldzlmzwqjpfalrmgzvxytnakvonwybyrmpdstxylyylquahxdesyoeflxuaowmwietrzhcydovcxzmwvaesvrgmuncwwybuexcvuqwythawoyqrpvkjopmnrukloeeukzeihhmyztuhkvoxhbzlcdiclacuhwcocklecpvpxetcomtejbiuxdgejxhruwshyvtsmtutxdmfnomlbgdfhhjbqrrfswbdchmocgljidpmevwhmwfghbrhrdvilvcshuhnyovggquxsjixfmvbctbfjyjpsqpikjhrmklgochnkpzshaeuelstfdrshqdkahniqvintyejtkbpcxmjqoaizfgabjzpbyfixjflhdsxesrbriiyzcuscioovtoegpllixfvdcqtfordjnkwgvsdpvriiwowcyhzoyjdgadfvxjwjqpyffcduvbswigrjkvcfsjeipygtavhssasophyjqlbmnxwegyglqblcaarkniingltxwjdxrtmaimfspmdgrnkukkfxfpkengoliarehwjvtxfgxzntduxdjkttjherhlxjofinyqpewhdxuknfppfvyqaswfsfcdhxtrmjgjssaedqykxkpmvwycqpkixnzdmospydtgvfhrmpdnhitdguaefrucqjwdrggtyzqnajpceoyodcvzxtqiljlsfiwfualsfldrjjncyayzznuzyqlzrcuzrlegdajoevjbxzquzujwppmajpojvxvdvisnxyvwpenrrkgqrrlvlffabcztzjvuepjvkkbjlwmutfhmgakgwlsdqldlygzizbqkoenqkdehkawctwyoulvjmphtogazamlyzbzmzqrfhkcviivgadlvsbjbfanqfjctqirvjdxzfiuvokryuxwmyjvawwjabiwshjcfohvtrsbqhxzxvicedhcxjeqcgghoydtowkofyegfzqquokskxpikqhbpfcjopiczgtwvzupqaqoiwngqbuvgmwczikirqrdkvznlvzabxopcqfncblnqhugwfzrjmcvsoubwahsrdjtqdkjhlrwwazonzmktcgxeejylpsevhghossmdkysbrpblmgcuvivuzixfioczrqobagvwmflqvyncnncytlwtartxxnqrzbdflmpzmgvxpjxyueqftmdojdekvzxlprkqrmgnwqgkirhvqqyknoyzpwzbcvteahdmifynjgjaieaurtqzcycsstthcbbfpgvludtburawtavtkytsrcmnnavhymgzoswdjpuxyfqxnylbtsbjotosrjlmveclsyikmnmrlgstmwzugqigcfrossfocaarywwzqkciyqpraoyymjxaxjablrybjnbjnfdslflkcjrooowyvhskbihtrswzxzrdbveeuuquflarbgeynlusyftmltrqhcowjaduphamghczyynxgiexviyqgdnzxgxttdwnlcfeqaluqfcrsmyswksqntdfwseybdcpyorbmoxfpveupqpyvvaatxdorrnddgtxqzneowjsaznumypnlooehgefiosqjnasueuadbmjzuyyqoocmoowfkqpcxktjrscbghwxzikowetxegrbrbmzwjuyhrkngpnpyjdmjcwyefqxxkjjijetuzarhluhnznsrsbchlikylflsaanryyqblzstmhlyfhasliiasrffzuwztzmdrsqkfnnggtkdrdtvzdfexbsuwqimrhdavdtcnknmktbuohlsscxvrdcjxtcfchentmshugglcuuknfzmgmicvzllytpfwqpabglnrsihuefgrlryhcviynkfhdjtxagxliigzhvhbayropnikarkkrklfptmnykswsxppsotyyqvmxozysnwctkdatcirssgkhgirfseyxijqmnnrnovavjimeanodhzrwchesfnnugsweheohatopkmnxgmvcyadviooyiaxdmaexrgvwmwakiqkzdvowibqeuwarsddwgupfqvajrmhpkatjzjcxnehfsetjcdewezvfpryoyvxhegadsuvcwbtpwirzapnhrewyouywqvsqhsqoyuwmgrprwtchxtvrzchyayhzvnvkrjdqqqeemljqrzpdzsdtqjksqctlvxdvyquscazlifqeupulfklomoldeybrxvdshgdtlkzribilqeugctkhvtosohmarlplcmpeocvfyjehfpdpnqfbsxynrbjgpaegiultvcvtsuuygwtzqhwzoxdrphjyctfvuxohxhhdvgqorxzhlzxpxtahwviouhlqjhwfoerbzzsoeccwnshezzlxjxjfqwayhrfiwbjvarqvtsowgkslfqzrfrqvjmlsrdvlvkkimikgjncdbokopqzhpkeifpexcksfmulqbtfgtbbmnljrjhzaiorzltrjisoxbrcvjowjzrjlklvfjwrwdukdcqqcfkyzkgpjrvyxmmeufjlbgskstyqkaekyrzvgxsvilggzlhsecuqapxxmuzgntmaqtjtzafgzrnipfrborhqyjtcbfyvpwnfstybhgybkqrfafhjokqtvauunxiqxzrnsjsgsbhpnbfptlrihwgguotuawlvqyppdbhjrqehtkffjdxkvqlwzgtfwvkqktibsxuiashdwdeskdgvgycuazmqvfnannldrnazjirwgdvmqsmaypbomgqivthyyxkryjqvfkkyascgebefkenlzbcpbiftzbdsinadxovxqrnofevwuqmjlfurgjymykdfztfaxyuhlzgraneragtijegqomgcpgojhuvpstpgwnvvxasjsforsojjtyuywsebnpkzhvnqxhftunzxjykguqyrmbuuyquoirsgupjegcoogocnfubnmtkpgjohcimgrxrtzusakfdxvkyayitryokwioklfvaxwcreiljhpywhgsojpkfetyuxvpbdfcjkzfrnhsgyuehstxytscdoxptzynbgxlmdfrqfbudtmtwcnooorhusgvlsbnxiottmqfqbyuuuoeapjgqnkmilfnkniiqtvrvzzfxmeekhbcqioztdndmjpodiznvghntoluazdxlaojzltdzquwjdhluvysxlygxwykmxlbbvoxyjsjujxwydisorlzmuvryjyqqrsmvcouxncounxdpmgnmomikvecyvgmuutbrxnuabxxxyzwyhwpbulhfunlqdgsxmzzkajksxeqkpxkbczhwgnmfdpkdncadknijbtyelmxhnnsxhrxpmcfoudtoetqsvicfnbxhhuihiehcfzqvawshpdramcvlomjfjahuizipdwfffxwrwfdiomyeexibqqgpybpuzsngiydmaxjnupalgbzmokmvlfmpgbixqqcxhhmhaatruihatgiqzlakajbucuzqjlocmtapnmlrbrdthytkfgqnngvvzblwwokuursysmcrole"
)

var tleInput2 string

func init() {
	tle2 := make([]byte, 10_000)
	for i := 0; i < 10_000; i++ {
		tle2[i] = 'a'
	}

	tleInput2 = string(tle2)
}