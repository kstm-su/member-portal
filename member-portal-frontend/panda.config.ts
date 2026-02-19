import { amber } from "./src/theme/colors/amber"
import { green } from "./src/theme/colors/green"
import { neutral } from "./src/theme/colors/neutral"
import { red } from "./src/theme/colors/red"
import { animationStyles } from "./src/theme/animation-styles"
import { conditions } from "./src/theme/conditions"
import { globalCss } from "./src/theme/global-css"
import { keyframes } from "./src/theme/keyframes"
import { layerStyles } from "./src/theme/layer-styles"
import { slotRecipes, recipes } from "./src/theme/recipes"
import { textStyles } from "./src/theme/text-styles"
import { colors } from "./src/theme/tokens/colors"
import { durations } from "./src/theme/tokens/durations"
import { shadows } from "./src/theme/tokens/shadows"
import { zIndex } from "./src/theme/tokens/z-index"
import { defineConfig } from "@pandacss/dev"

export default defineConfig({
  preflight: true,
  include: ["./src/**/*.{js,jsx,ts,tsx}"],
  exclude: [],
  outdir: "styled-system",
  jsxFramework: "react",

  conditions,

  theme: {
    extend: {
      semanticTokens: {
        colors: {
          error: {
            value: {
              _light: "{colors.red.9}",
              _dark: "{colors.red.9}"
            }
          },

          amber: amber,
          gray: neutral,
          red: red,
          green: green
        },

        shadows: shadows,

        radii: {
          l1: {
            value: "{radii.xs}"
          },

          l2: {
            value: "{radii.sm}"
          },

          l3: {
            value: "{radii.md}"
          }
        }
      },

      animationStyles: animationStyles,
      recipes: recipes,
      slotRecipes: slotRecipes,
      keyframes: keyframes,
      layerStyles: layerStyles,
      textStyles: textStyles,

      tokens: {
        colors: colors,
        durations: durations,
        zIndex: zIndex
      }
    }
  },

  globalCss: globalCss
})
