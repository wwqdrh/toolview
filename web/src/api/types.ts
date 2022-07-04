interface EtcdList {
  name: string
  value: string
}

interface EtcdConfig {
  endpoints: string
  username: string
  password: string
}

interface RedisConfig {
  endpoint: string
  password: string
}

interface APIResponse {
  code: number
  msg: string
  description: string
  data: object
}

export type { EtcdList, EtcdConfig, APIResponse, RedisConfig }
