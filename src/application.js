require('./util/fetch')

import React from 'react'
import ReactDOM from 'react-dom'

var el = document.createElement('div')
ReactDOM.render(React.createElement(component, {}, null), el)
document.body.appendChild(el)
