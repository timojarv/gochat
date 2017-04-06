module.exports = {
    context: __dirname + "/src",
    entry: "./index.js",
    output: {
        path: __dirname + "/public",
        filename: "bundle.js",
        publicPath: "/"
    },
    module: {
        loaders: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['env', 'react', 'stage-0']
                    }
                }
            },
            {
                test: /\.s?css$/,
                loader: "style-loader!css-loader!sass-loader"
            }
        ]
    }
}