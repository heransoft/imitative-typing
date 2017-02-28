(function () {
    var path = "/path/aa";
    it.d.i(path);
    var fileCount = it.d.fc(path);
    it.g={};
    it.g.uuidMap = {};
    it.g.uuidArray = [];
    it.s.f["get_uuid"]=function () {
        var lineCount = it.f.lc();
        for(var i = 0;i<lineCount;i++){
            var lineContent = it.l.s(i);
            var reg = new RegExp('\"__uuid__\": \"([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})\"');
            var uuid = lineContent.match(reg);
            if(uuid != null && uuid.length == 2){
                if(it.g.uuidMap[uuid[1]] != 1){
                    it.g.uuidMap[uuid[1]] = 1;
                    it.g.uuidArray.push(uuid[1]);
                }
            }
        }
    };

     for (var i = 0; i < fileCount; i++) {
        var filename = it.d.fn(i);
         it.f.u(filename, JSON.stringify(["get_uuid"]));
    }

    it.g.uuidArray.sort();
    it.s.f["gen_uuid_text"] = function () {
        for (var i in it.g.uuidArray) {
            it.l.i(it.g.uuidArray[i] + '\n');
        }
    };
    it.f.i("","uuid_text.txt",JSON.stringify(["gen_uuid_text"]));
})();