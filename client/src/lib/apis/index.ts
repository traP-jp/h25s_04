import { Apis, Configuration } from './generated'

const apis = new Apis(new Configuration({ basePath: '/api/v1' }))

export default apis
export * from './generated'
