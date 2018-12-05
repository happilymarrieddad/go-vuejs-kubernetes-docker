const path = require('path')

module.exports = {
    configureWebpack: {
        resolve: {
            extensions: ['.js','.vue','.json','.scss'],
            alias: {
                'styles': path.resolve(__dirname, 'src/assets/scss')
            }
        }
    }
}