const {execFileSync} = require("child_process")


class DotenvParser{
    constructor(binarypath = "./universal-dotenv"){
        this.binarypath = binarypath
    }

    read(state){
        try{
            const output = execFileSync(this.binarypath,[state],{encoding : "utf-8"})

            try{
                return JSON.parse(output)
            }catch{
                return output.trim()
            }
        }catch(err){
            throw new Error(`Failed to run the binding binary ${err.message}`)
        }
    }
}

const reader = new DotenvParser();
const env = reader.read("true");

console.log(env);