import * as path from "path";
import * as fs from "fs";

import * as Bundler from "parcel-bundler";

const entry = process.argv[2];
const target = process.argv[3];

const fatal = (...messages) => {
    console.error(...messages);
    process.exit(1);
};

if (typeof entry !== "string") {
    fatal("entry not defined (first argument)");
}

if (typeof target !== "string") {
    fatal("target not defined (second argument)");
}

process.env.NODE_ENV = "production";

const bundler = new Bundler(path.join(__dirname, entry), {
    outDir: target,
    sourceMaps: false,
    logLevel: 1,
});

bundler.bundle().then((rootBundle) => {
    const bundles = scrapeBundles(rootBundle);

    if (!bundles.html || bundles.html.length !== 1) {
        fatal("must have exactly one html bundle");
    }
    if (!bundles.js || bundles.js.length !== 1) {
        fatal("must have exactly one js bundle");
    }
    if (!bundles.css || bundles.css.length !== 1) {
        fatal("must have exactly one css bundle");
    }

    let html = fs.readFileSync(bundles.html[0], "utf8");

    html = rename(html, bundles.js[0], "script.js");
    html = rename(html, bundles.css[0], "style.css");

    fs.writeFileSync(bundles.html[0], html);
}).catch(fatal);

const scrapeBundles = (bundle) => {
    const paths: {[type: string]: string[]} = {};

    const traverse = (bundle) => {
        const {type} = bundle;
        if (!paths[type]) {
            paths[type] = [];
        }
        paths[type].push(bundle.name);

        bundle.childBundles.forEach((child) => {
            traverse(child);
        });
    };
    traverse(bundle);

    return paths;
};

const rename = (template, from, to) => {
    const filepath = path.dirname(from);
    const filename = path.basename(from);
    fs.renameSync(from, path.join(filepath, to));
    return template.replace(filename, to);
};
