package gowd

import (
    "testing"
    "gotest.tools/assert"
    "encoding/json"
    _ "fmt"
)


func TestDataModelUnmarshal(t *testing.T) {
    var data struct {
        Entities map[string]Entity  `json:"entities"`
    }
    
    err := json.Unmarshal([]byte(entitydata_Q570349_oldid_877006223), &data)
    if err != nil {
        t.Log("Ummarshalling Error")
        t.Fatal(err)
    }
    
    Q570349 := data.Entities["Q570349"]
    
    assert.Equal(t, Q570349.PageId.Int64(), int64(536525))
    assert.Equal(t, Q570349.Title, "Q570349")
    assert.Equal(t, Q570349.Type, EntityType_Item)
    assert.Equal(t, string(Q570349.Type), "item")
    assert.Equal(t, Q570349.Lastrevid.Int64(), int64(877006223))
    assert.Equal(t, Q570349.Id, "Q570349")

    assert.Equal(t, Q570349.Labels["zh"].Language,   "zh")
    assert.Equal(t, Q570349.Labels["zh"].Value,      "三峽區")
    assert.Equal(t, Q570349.Labels["en"].Language,   "en")
    assert.Equal(t, Q570349.Labels["en"].Value,      "Sanxia District")
    assert.Equal(t, Q570349.Aliases["zh-tw"][0].Language,   "zh-tw")
    assert.Equal(t, Q570349.Aliases["zh-tw"][0].Value,      "三角湧")
    assert.Equal(t, Q570349.Aliases["zh-tw"][1].Value,      "三角躅")
    assert.Equal(t, Q570349.Aliases["zh-tw"][2].Value,      "三峽鎮")

    mainSnak := &(Q570349.Claims["P17"][0].MainSnak)
    assert.Equal(t, mainSnak.SnakType, "value")
    assert.Equal(t, mainSnak.Property, "P17")
    assert.Equal(t, mainSnak.DataValue.Type(), "wikibase-entityid")
    assert.Equal(t, mainSnak.DataValue.(*WikibaseEntityId).EntityType, "item")
    assert.Equal(t, mainSnak.DataValue.(*WikibaseEntityId).NumericId.Int64(), int64(865))
    assert.Equal(t, mainSnak.DataValue.(*WikibaseEntityId).Value, "Q865")

    // TODO: Add test for qualifiers
}

var ent_1 string = `
    {
        "entities":{
            "Q123":{
                "pageid": 9487,
                "type": "item",
                "id": "Q123"
            }
        }
    }
`

// url: <https://www.wikidata.org/wiki/Special:EntityData?id=Q570349&oldid=877006223&format=json>
var entitydata_Q570349_oldid_877006223 string = `{"entities":{"Q570349":{"pageid":536525,"ns":0,"title":"Q570349","lastrevid":877006223,"modified":"2019-03-08T01:43:49Z","type":"item","id":"Q570349","labels":{"zh":{"language":"zh","value":"\u4e09\u5cfd\u5340"},"en":{"language":"en","value":"Sanxia District"},"en-ca":{"language":"en-ca","value":"Sanxia District"},"en-gb":{"language":"en-gb","value":"Sanxia District"},"fr":{"language":"fr","value":"Sanxia"},"hak":{"language":"hak","value":"S\u00e2m-hia\u030dp-kh\u00ee"},"it":{"language":"it","value":"Sanxia"},"ja":{"language":"ja","value":"\u4e09\u5cfd\u533a"},"ko":{"language":"ko","value":"\uc2fc\uc0e4\uad6c"},"vi":{"language":"vi","value":"Tam H\u1ea1p, T\u00e2n B\u1eafc"},"de":{"language":"de","value":"Sanxia"},"nan":{"language":"nan","value":"Sam-kiap-khu"},"lzh":{"language":"lzh","value":"\u4e09\u5cfd\u5340"},"lt":{"language":"lt","value":"Sansija"},"nl":{"language":"nl","value":"Sanxia District"},"uk":{"language":"uk","value":"\u0421\u0430\u043d\u044c\u0441\u044f"},"pl":{"language":"pl","value":"Sanxia"},"zh-hant":{"language":"zh-hant","value":"\u4e09\u5cfd\u5340"},"zh-tw":{"language":"zh-tw","value":"\u4e09\u5cfd\u5340"}},"descriptions":{"zh":{"language":"zh","value":"\u4e09\u5cfd\u5340\u70ba\u53f0\u7063\u65b0\u5317\u5e02\u7684\u5e02\u8f44\u5340\u4e4b\u4e00"},"en":{"language":"en","value":"district in New Taipei City"},"de":{"language":"de","value":"Bezirk der Stadt Neu-Taipeh, Republik China"},"zh-hant":{"language":"zh-hant","value":"\u4e09\u5cfd\u5340\u70ba\u53f0\u7063\u65b0\u5317\u5e02\u7684\u5e02\u8f44\u5340\u4e4b\u4e00"},"zh-tw":{"language":"zh-tw","value":"\u4e09\u5cfd\u5340\u70ba\u53f0\u7063\u65b0\u5317\u5e02\u7684\u5e02\u8f44\u5340\u4e4b\u4e00"}},"aliases":{"zh-tw":[{"language":"zh-tw","value":"\u4e09\u89d2\u6e67"},{"language":"zh-tw","value":"\u4e09\u89d2\u8e85"},{"language":"zh-tw","value":"\u4e09\u5cfd\u93ae"}],"zh":[{"language":"zh","value":"\u4e09\u89d2\u8e85"},{"language":"zh","value":"\u4e09\u89d2\u6e67"},{"language":"zh","value":"\u4e09\u5cfd\u93ae"}],"zh-hant":[{"language":"zh-hant","value":"\u4e09\u89d2\u8e85"},{"language":"zh-hant","value":"\u4e09\u89d2\u6e67"},{"language":"zh-hant","value":"\u4e09\u5cfd\u93ae"}],"en":[{"language":"en","value":"Sansia District"}]},"claims":{"P17":[{"mainsnak":{"snaktype":"value","property":"P17","datavalue":{"value":{"entity-type":"item","numeric-id":865,"id":"Q865"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"q570349$0F2973D6-6E41-4E61-9D24-EDBB472B8B2F","rank":"normal"}],"P131":[{"mainsnak":{"snaktype":"value","property":"P131","datavalue":{"value":{"entity-type":"item","numeric-id":244898,"id":"Q244898"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$177acd8f-432e-31d3-fa60-28ae7c392b3a","rank":"normal"}],"P373":[{"mainsnak":{"snaktype":"value","property":"P373","datavalue":{"value":"Sanxia District, New Taipei","type":"string"},"datatype":"string"},"type":"statement","id":"q570349$B66735E1-FA14-4627-BF22-D3877A0F96F9","rank":"normal"}],"P625":[{"mainsnak":{"snaktype":"value","property":"P625","datavalue":{"value":{"latitude":24.9347007,"longitude":121.3707733,"altitude":null,"precision":2.77777777778e-6,"globe":"http://www.wikidata.org/entity/Q2"},"type":"globecoordinate"},"datatype":"globe-coordinate"},"type":"statement","id":"Q570349$847E89CF-A1E6-4BFC-AADA-FC2D9366B27F","rank":"normal","references":[{"hash":"d4bd87b862b12d99d26e86472d44f26858dee639","snaks":{"P143":[{"snaktype":"value","property":"P143","datavalue":{"value":{"entity-type":"item","numeric-id":8447,"id":"Q8447"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}]},"snaks-order":["P143"]}]}],"P646":[{"mainsnak":{"snaktype":"value","property":"P646","datavalue":{"value":"/m/05c0b9w","type":"string"},"datatype":"external-id"},"type":"statement","id":"Q570349$F2E988B2-D8CB-485F-8B17-0B7CEF70195E","rank":"normal","references":[{"hash":"2b00cb481cddcac7623114367489b5c194901c4a","snaks":{"P248":[{"snaktype":"value","property":"P248","datavalue":{"value":{"entity-type":"item","numeric-id":15241312,"id":"Q15241312"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}],"P577":[{"snaktype":"value","property":"P577","datavalue":{"value":{"time":"+2013-10-28T00:00:00Z","timezone":0,"before":0,"after":0,"precision":11,"calendarmodel":"http://www.wikidata.org/entity/Q1985727"},"type":"time"},"datatype":"time"}]},"snaks-order":["P248","P577"]}]}],"P31":[{"mainsnak":{"snaktype":"value","property":"P31","datavalue":{"value":{"entity-type":"item","numeric-id":705296,"id":"Q705296"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","qualifiers":{"P580":[{"snaktype":"value","property":"P580","hash":"f40f7a3b9b28e53e1034a3e8568f93a60b52188e","datavalue":{"value":{"time":"+2010-00-00T00:00:00Z","timezone":0,"before":0,"after":0,"precision":9,"calendarmodel":"http://www.wikidata.org/entity/Q1985727"},"type":"time"},"datatype":"time"}]},"qualifiers-order":["P580"],"id":"Q570349$6EEE4D9F-A618-4C71-A377-256F594AFED3","rank":"normal"},{"mainsnak":{"snaktype":"value","property":"P31","datavalue":{"value":{"entity-type":"item","numeric-id":12039044,"id":"Q12039044"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","qualifiers":{"P582":[{"snaktype":"value","property":"P582","hash":"6f4c37357db10714dc6c7a0ce30a4621b3ce0305","datavalue":{"value":{"time":"+2010-00-00T00:00:00Z","timezone":0,"before":0,"after":0,"precision":9,"calendarmodel":"http://www.wikidata.org/entity/Q1985727"},"type":"time"},"datatype":"time"}]},"qualifiers-order":["P582"],"id":"Q570349$eec517c5-42f6-dd55-7a40-ac0e4e6f0ac3","rank":"normal"}],"P1566":[{"mainsnak":{"snaktype":"value","property":"P1566","datavalue":{"value":"1670106","type":"string"},"datatype":"external-id"},"type":"statement","id":"Q570349$33A8819A-2D21-4B40-95E2-886F922E4218","rank":"normal","references":[{"hash":"64133510dcdf15e7943de41e4835c673fc5d6fe4","snaks":{"P143":[{"snaktype":"value","property":"P143","datavalue":{"value":{"entity-type":"item","numeric-id":830106,"id":"Q830106"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}]},"snaks-order":["P143"]}]}],"P910":[{"mainsnak":{"snaktype":"value","property":"P910","datavalue":{"value":{"entity-type":"item","numeric-id":9446659,"id":"Q9446659"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$A3F5622C-A512-4DC9-A749-4195D70364F9","rank":"normal","references":[{"hash":"0ee3b3ba1c958f4c3dcba7ed8091fe4b57311348","snaks":{"P143":[{"snaktype":"value","property":"P143","datavalue":{"value":{"entity-type":"item","numeric-id":30239,"id":"Q30239"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}]},"snaks-order":["P143"]}]}],"P18":[{"mainsnak":{"snaktype":"value","property":"P18","datavalue":{"value":"SanxiaPaifangTaiwan.jpg","type":"string"},"datatype":"commonsMedia"},"type":"statement","id":"Q570349$533CD4AC-E839-46AB-BDDA-9F1C9E8CD6A6","rank":"normal","references":[{"hash":"fa278ebfc458360e5aed63d5058cca83c46134f1","snaks":{"P143":[{"snaktype":"value","property":"P143","datavalue":{"value":{"entity-type":"item","numeric-id":328,"id":"Q328"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}]},"snaks-order":["P143"]}]}],"P856":[{"mainsnak":{"snaktype":"value","property":"P856","datavalue":{"value":"http://www.sanshia.tpc.gov.tw/","type":"string"},"datatype":"url"},"type":"statement","id":"Q570349$BCF8EB8F-EEF8-4EBC-B0EF-E506B3417C38","rank":"normal","references":[{"hash":"d5847b9b6032aa8b13dae3c2dfd9ed5d114d21b3","snaks":{"P143":[{"snaktype":"value","property":"P143","datavalue":{"value":{"entity-type":"item","numeric-id":11920,"id":"Q11920"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}]},"snaks-order":["P143"]}]}],"P948":[{"mainsnak":{"snaktype":"value","property":"P948","datavalue":{"value":"Sanxia (Taiwan) wikivoyage banner.jpg","type":"string"},"datatype":"commonsMedia"},"type":"statement","id":"Q570349$6F856BD6-0A57-4C61-AD53-A1D55ACF0D22","rank":"normal","references":[{"hash":"b463db83d7bee30412bc8e17b63ac4669393fb50","snaks":{"P143":[{"snaktype":"value","property":"P143","datavalue":{"value":{"entity-type":"item","numeric-id":24238024,"id":"Q24238024"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}]},"snaks-order":["P143"]}]}],"P571":[{"mainsnak":{"snaktype":"value","property":"P571","datavalue":{"value":{"time":"+2010-00-00T00:00:00Z","timezone":0,"before":0,"after":0,"precision":9,"calendarmodel":"http://www.wikidata.org/entity/Q1985727"},"type":"time"},"datatype":"time"},"type":"statement","id":"Q570349$ef882b5d-4f84-cffc-3cca-71def05f3466","rank":"normal"}],"P2046":[{"mainsnak":{"snaktype":"value","property":"P2046","datavalue":{"value":{"amount":"+191","unit":"http://www.wikidata.org/entity/Q712226"},"type":"quantity"},"datatype":"quantity"},"type":"statement","id":"Q570349$03aec289-4a79-810f-d3fb-3a85d82f91a8","rank":"normal"}],"P421":[{"mainsnak":{"snaktype":"value","property":"P421","datavalue":{"value":{"entity-type":"item","numeric-id":712168,"id":"Q712168"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$5464a19d-4f49-df3d-39d7-10d4574f08fa","rank":"normal"}],"P281":[{"mainsnak":{"snaktype":"value","property":"P281","datavalue":{"value":"237","type":"string"},"datatype":"string"},"type":"statement","id":"Q570349$767832cf-4cde-0fc8-ed33-39dcdcd8134b","rank":"normal"}],"P242":[{"mainsnak":{"snaktype":"value","property":"P242","datavalue":{"value":"New Taipei-Sanxia.svg","type":"string"},"datatype":"commonsMedia"},"type":"statement","id":"Q570349$9D4C1ED7-D62F-49C5-A99A-C18883EA07D9","rank":"normal","references":[{"hash":"4dd821e0419d2584ea30fac1ea9fc62cd883492f","snaks":{"P143":[{"snaktype":"value","property":"P143","datavalue":{"value":{"entity-type":"item","numeric-id":328,"id":"Q328"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}],"P4656":[{"snaktype":"value","property":"P4656","datavalue":{"value":"https://en.wikipedia.org/w/index.php?title=Sanxia_District&oldid=805158835","type":"string"},"datatype":"url"}]},"snaks-order":["P143","P4656"]}]}],"P1082":[{"mainsnak":{"snaktype":"value","property":"P1082","datavalue":{"value":{"amount":"+115670","unit":"1"},"type":"quantity"},"datatype":"quantity"},"type":"statement","qualifiers":{"P585":[{"snaktype":"value","property":"P585","hash":"079c86e11516c98c01ab2605ed797b3a776460c5","datavalue":{"value":{"time":"+2018-09-00T00:00:00Z","timezone":0,"before":0,"after":0,"precision":10,"calendarmodel":"http://www.wikidata.org/entity/Q1985727"},"type":"time"},"datatype":"time"}]},"qualifiers-order":["P585"],"id":"Q570349$de560f3a-4cc6-bb5f-69af-e4d101f3fef4","rank":"normal"}],"P47":[{"mainsnak":{"snaktype":"value","property":"P47","datavalue":{"value":{"entity-type":"item","numeric-id":570493,"id":"Q570493"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$28164e71-439e-ee77-8716-6accfe543aca","rank":"normal"},{"mainsnak":{"snaktype":"value","property":"P47","datavalue":{"value":{"entity-type":"item","numeric-id":516297,"id":"Q516297"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$1795c7af-4a60-f10b-73bd-5a1ad5893b70","rank":"normal"},{"mainsnak":{"snaktype":"value","property":"P47","datavalue":{"value":{"entity-type":"item","numeric-id":372120,"id":"Q372120"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$2ed03da5-4274-78bb-f732-d39e5bff75aa","rank":"normal"},{"mainsnak":{"snaktype":"value","property":"P47","datavalue":{"value":{"entity-type":"item","numeric-id":340192,"id":"Q340192"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$fde24f59-452f-96dd-b957-d65e7588c952","rank":"normal"},{"mainsnak":{"snaktype":"value","property":"P47","datavalue":{"value":{"entity-type":"item","numeric-id":32728,"id":"Q32728"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$8686a108-4206-4c30-c03a-f6575c9a9d2f","rank":"normal"},{"mainsnak":{"snaktype":"value","property":"P47","datavalue":{"value":{"entity-type":"item","numeric-id":706455,"id":"Q706455"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$1eaf3ae4-4464-f70c-3deb-16f87fc6d0e7","rank":"normal"},{"mainsnak":{"snaktype":"value","property":"P47","datavalue":{"value":{"entity-type":"item","numeric-id":718377,"id":"Q718377"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$8d890309-4636-b7bf-b317-64d8ff8ef874","rank":"normal"}],"P1365":[{"mainsnak":{"snaktype":"value","property":"P1365","datavalue":{"value":{"entity-type":"item","numeric-id":10865954,"id":"Q10865954"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$7214f07f-472b-2974-f109-b5a913c7fa40","rank":"normal"}],"P402":[{"mainsnak":{"snaktype":"value","property":"P402","datavalue":{"value":"2858710","type":"string"},"datatype":"external-id"},"type":"statement","id":"Q570349$721FE8C8-AACA-4FAB-A143-F110951A2BD2","rank":"normal"}],"P797":[{"mainsnak":{"snaktype":"value","property":"P797","datavalue":{"value":{"entity-type":"item","numeric-id":61788154,"id":"Q61788154"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$5aff6b84-45a2-5a4c-2450-18d188a742f7","rank":"normal"}],"P1313":[{"mainsnak":{"snaktype":"value","property":"P1313","datavalue":{"value":{"entity-type":"item","numeric-id":61788436,"id":"Q61788436"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},"type":"statement","id":"Q570349$6780f517-4fbf-50db-725c-aaa9a6ad2cde","rank":"normal"}],"P3722":[{"mainsnak":{"snaktype":"value","property":"P3722","datavalue":{"value":"Maps of Sanxia District, New Taipei","type":"string"},"datatype":"string"},"type":"statement","id":"Q570349$a35ca0ca-1fa8-4e9d-84ca-8241d987b40d","rank":"normal"}]},"sitelinks":{"dewiki":{"site":"dewiki","title":"Sanxia (Neu-Taipeh)","badges":[],"url":"https://de.wikipedia.org/wiki/Sanxia_(Neu-Taipeh)"},"enwiki":{"site":"enwiki","title":"Sanxia District","badges":[],"url":"https://en.wikipedia.org/wiki/Sanxia_District"},"frwiki":{"site":"frwiki","title":"Sanxia","badges":[],"url":"https://fr.wikipedia.org/wiki/Sanxia"},"hakwiki":{"site":"hakwiki","title":"S\u00e2m-hia\u030dp-kh\u00ee","badges":[],"url":"https://hak.wikipedia.org/wiki/S%C3%A2m-hia%CC%8Dp-kh%C3%AE"},"itwiki":{"site":"itwiki","title":"Sanxia","badges":[],"url":"https://it.wikipedia.org/wiki/Sanxia"},"jawiki":{"site":"jawiki","title":"\u4e09\u5ce1\u533a","badges":[],"url":"https://ja.wikipedia.org/wiki/%E4%B8%89%E5%B3%A1%E5%8C%BA"},"kowiki":{"site":"kowiki","title":"\uc2fc\uc0e4\uad6c","badges":[],"url":"https://ko.wikipedia.org/wiki/%EC%8B%BC%EC%83%A4%EA%B5%AC"},"ltwiki":{"site":"ltwiki","title":"Sansija","badges":[],"url":"https://lt.wikipedia.org/wiki/Sansija"},"plwiki":{"site":"plwiki","title":"Sanxia","badges":[],"url":"https://pl.wikipedia.org/wiki/Sanxia"},"ukwiki":{"site":"ukwiki","title":"\u0421\u0430\u043d\u044c\u0441\u044f (\u0421\u0456\u043d\u044c\u0431\u0435\u0439)","badges":[],"url":"https://uk.wikipedia.org/wiki/%D0%A1%D0%B0%D0%BD%D1%8C%D1%81%D1%8F_(%D0%A1%D1%96%D0%BD%D1%8C%D0%B1%D0%B5%D0%B9)"},"viwiki":{"site":"viwiki","title":"Tam H\u1ea1p, T\u00e2n B\u1eafc","badges":[],"url":"https://vi.wikipedia.org/wiki/Tam_H%E1%BA%A1p,_T%C3%A2n_B%E1%BA%AFc"},"zh_classicalwiki":{"site":"zh_classicalwiki","title":"\u4e09\u5cfd\u5340","badges":[],"url":"https://zh-classical.wikipedia.org/wiki/%E4%B8%89%E5%B3%BD%E5%8D%80"},"zh_min_nanwiki":{"site":"zh_min_nanwiki","title":"Sam-kiap-khu","badges":[],"url":"https://zh-min-nan.wikipedia.org/wiki/Sam-kiap-khu"},"zhwiki":{"site":"zhwiki","title":"\u4e09\u5cfd\u5340","badges":[],"url":"https://zh.wikipedia.org/wiki/%E4%B8%89%E5%B3%BD%E5%8D%80"},"zhwikivoyage":{"site":"zhwikivoyage","title":"\u65b0\u5317/\u4e09\u5cfd","badges":[],"url":"https://zh.wikivoyage.org/wiki/%E6%96%B0%E5%8C%97/%E4%B8%89%E5%B3%BD"}}}}}`