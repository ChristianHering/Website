module.exports = {
    productionSourceMap: false, //Disabling this in production damages SEO
    outputDir: "./../../cdn",
    indexPath: "./../utils/templates/adminDashboard.html",
    integrity: false,
    filenameHashing: false, //Enable this for better user cache control
    publicPath: "http://cdn.ChristianHering.com/"
}
