const taskQueue = []

function deploy(env, branch){
    taskQueue.push(new Task(env, branch))
}

function removeTask(targetId){
    taskQueue.splice(taskQueue.findIndex(({id}) => id === targetId), 1)
}