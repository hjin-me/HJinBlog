module.exports = function(disqus_shortname, disqus_identifier) {

    window.disqus_shortname = disqus_shortname
    window.disqus_identifier = disqus_identifier
    var dsq = document.createElement('script')
    dsq.type = 'text/javascript'
    dsq.async = true
    dsq.src = '//' + disqus_shortname + '.disqus.com/embed.js'
    document.getElementsByTagName('head')[0].appendChild(dsq)
}