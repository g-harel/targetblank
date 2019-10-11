import {normalize} from "csstips";
import {cssRule} from "typestyle";

import {color, font, size} from "./theme";

normalize();

cssRule("html", {
    color: color.textPrimary,
    fontSize: size.normal,
});

cssRule("body", {
    backgroundColor: color.backgroundSecondary,
    fontSize: "100%",
    height: "100vh",
});

cssRule("*", {
    boxSizing: "border-box",
    fontFamily: font.normal,
});

cssRule("a", {
    textDecoration: "none",
});
