function q(selector){
    const ele = document.querySelector(selector)

    return {
        ele,
        css(propName, value){
            ele.style[propName] = `${value}`
        }
    }
}