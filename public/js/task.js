const TASK_NAME_MAP = ['测试环境1', '测试环境2', '本地环境']

class Task {
    constructor(env, branch){
        this.name = `${TASK_NAME_MAP[env]}`
        this.env = env
        this.branch = branch
        this.id = Date.now()
    }
}