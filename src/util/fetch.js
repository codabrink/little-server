import merge       from 'deepmerge'
import {stringify} from 'query-string'
import                  'whatwg-fetch'

var el    = document.querySelector('meta[name=csrf-token]')
var token = el ? el.getAttribute('content') : ''

window.CSRFFetch = function(url, options = {}) {
  options.body = JSON.stringify(options.body)
  return fetch(
    `${url}?${stringify(options.params)}`,
    merge({
      headers: {
        'X-CSRF-Token':     token,
        'X-Requested-With': 'XMLHttpRequest'
      },
      credentials: 'same-origin'
    }, options))
}
window.jCSRFFetch = function(url, options = {}) {
  return new Promise(function(resolve, reject) {
    window
      .CSRFFetch(
        url,
        merge({
          headers: {
            'Content-Type': 'application/json',
            'Accept':       'application/json'
          }
        }, options))
      .then(function(response) {
        response.text().then(function(text) {
          try {
            resolve({
              status: response.status,
              data: JSON.parse(text)
            })
          } catch(err) {
            resolve({
              status: response.status,
              data: {}
            })
          }
        })
      })
  })
}

window.get   = function(url, options = {}) { return window.jCSRFFetch(url, merge(options, {method: 'GET'})) }
window.post  = function(url, options = {}) { return window.jCSRFFetch(url, merge(options, {method: 'POST'})) }
window.patch = function(url, options = {}) { return window.jCSRFFetch(url, merge(options, {method: 'PATCH'})) }
window.del   = function(url, options = {}) { return window.jCSRFFetch(url, merge(options, {method: 'DELETE'})) }
