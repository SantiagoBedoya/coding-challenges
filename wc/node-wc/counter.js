
const countBytes = (input) => {
    return Buffer.byteLength(input, 'utf-8')
}

const countLines = (input) => {
    return input.split('\n').length;
}

const countWords = (input) => {
    return input.trim().split(/\s+/).length;
}

const countCharacters = (input) => {
    return input.length
}

module.exports = {
    processInput: (input, options) => {
        let response = "";
        
        if (options.c) {
            const bytes = countBytes(input)
            response += `\t${bytes}`
        }

        if (options.l) {
            const lines = countLines(input)
            response += `\t${lines}`  
        }

        if (options.w) {
            const words = countWords(input)
            response += `\t${words}`
        }

        if (options.m) {
            const characters = countCharacters(input)
            response += `\t${characters}`
        }

        if (!options.c && !options.l && !options.w && !options.m) {
            const bytes = countBytes(input)
            const lines = countLines(input)
            const words = countWords(input)
            const characters = countCharacters(input)
            response = `\t${bytes}\t${lines}\t${words}\t${characters}`
        }

        if (!options.isFromStdin) {
            response += `\t${options.filePath}`
        }

        console.log(response)
    }
}