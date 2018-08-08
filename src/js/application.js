require('./util/fetch')

import React     from 'react'
import ReactDOM  from 'react-dom'

import Container from './components/container.jsx'

var el = document.createElement('div')
ReactDOM.render(React.createElement(Container, {}, null), el)
document.body.appendChild(el)
