module.exports = {
    productionSourceMap: false, //Disabling this in production damages SEO
    outputDir: "./../../cdn",
    indexPath: "./../utils/templates/adminIndex.html",
    integrity: true,
    filenameHashing: true //Enable this for better user cache control
}
