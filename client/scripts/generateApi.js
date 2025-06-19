import fs from 'fs/promises'
import path from 'path'
import { exec } from 'child_process'
import { promisify } from 'util'
import { addApis } from './addApis.js'

const __dirname = import.meta.dirname

const execPromise = promisify(exec)

const SWAGGER_PATH = 'schema/openapi.yaml'
const GENERATED_DIR = 'src/lib/apis/generated'

const generateCmd = [
    'docker run --rm -v "$(dirname "$PWD"):/local" -u $(id -u) openapitools/openapi-generator-cli:v7.5.0',
    'generate',
    '-g',
    'typescript-axios',
    '-i',
    `/local/${SWAGGER_PATH}`,
    '-o',
    `/local/client/${GENERATED_DIR}`,
]

;(async () => {
    await fs.mkdir(path.resolve(__dirname, '../', GENERATED_DIR), {
        recursive: true,
    })

    await execPromise(generateCmd.join(' '))

    // generate Apis class
    await addApis(GENERATED_DIR)
})()
