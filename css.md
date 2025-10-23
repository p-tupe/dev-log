---
modified: Sat Nov 26 13:21:39 EST 2022
---
# css

- [MDN Docs](https://developer.mozilla.org/en-US/docs/Learn/CSS)

- [freeCodeCamp youtube link](https://www.youtube.com/watch?v=a_iQb1lnAEQ)

## Selectors

```css
<selector > {
  <property>: <value>;
}
```

- Element Selectors `h1, body`
- ID Selectors `#id-name`
- Class Selectors `.class-name`
- Wildcard Selectors `[<attr>*="<str>"]`, `^` = starts with, `$` = ends with
- Universal Selector (type of Wildcard) `*`

## Pseudo-Classes

`<selector>:<pseudo-class> {...}`

- hover
- visited
- checked
- first-child
- last-child
- nth-child - `<selector>:nth-child(<n> | even | odd)`

## Pesudo-Elements

`<selector>::<pseudo-element> {...}`

- first-letter
- first-line
- before, after
  - `{ content: "some content to add without changing html" }`

## Combinators

`<selector> <combinator> <selector2> {...}`

- direct child `>`
- all descendants `<literal space>`
- adjacent siblings `+`
- general siblings (all) `~`

## More on Wildcard Selectors

- To select all classes containing `str` in name:

  ```css
  [class*="str"] ...
  ```

- To select all classes that start with `strt` or end with `end` use:

  ```css
  [class^="strt"] ...
  [class$="end"] ...
  ```

## Tips

- To center image: `margin: 0 auto`

- To center text: `text-align: center`

- Group duplicates: `h1, h2, .header-class, #header-id { ... }`

## Sizing Units

- Absolute units: px, pt, cm, mm, in

- Relative units: em, rem, vw, vh, %

- Standard html size is around 16px

- Use relative units for responsive UI

## Colors

- Build-in names `white`, `blue`, `black`, `red`
- RGB(A) (Red, Blue, Green, Alpha) `rgb(x, y, z)`
- Hexadecimal `#000000` -> `#FFFFFF` #RRBBGG
- HSL (Hue, Saturation%, Lightness%)

- [Adobe Color Wheel](https://www.color.adobe.com/create/color-wheel)

## Typography

- font-family
- font-weight
- text-align
- text-shadow
- text-transform
- text-indent
- line-height
- letter-spacing
- word-spacing
- white-space
- background-color (not specific to typography)

## Box-model

> content, padding, border, margin.

## Layouts

### Flexbox

- Parent:

  - `display: flex`
  - `flex-direction: row (default on web) | column (default on react-native)`
  - `justify-content: center | ...` (main axis)
  - `align-items: center | flex-start | ...` (cross axis)
  - `flex-wrap: wrap | nowrap`

- Children:

  - `flex-grow: <number>` amount of space an item will take
  - `flex-shrink: <number>` same as above
  - `flex-basis: <value>` Initial size
  - Shorthand for grow, shrink and basis: `flex: <num> <num> <val>`
  - `align-self: center | ...`

### Grid

## Frameworks

- [Bulma](./bulma.md)
- Bootstrap
- Ant
- Sementic

### Tips n Tricks

- Center Div

  ```html
  <div class="outer">This text should be in center!</div>

  <style>
    .outer {
      display: flex /* grid */;
      justify-content: center;
      align-items: center;
      height: 100%;
    }
  </style>
  ```

- Change svg color using css filter

[Source](https://stackoverflow.com/questions/22252472/how-to-change-the-color-of-an-svg-element)

1. Add the SVG image using an `<img>` tag.

   ```html
   <img src="dotted-arrow.svg" class="filter-green" />
   ```

2. To filter to a specific color, use the following [Codepen](https://codepen.io/sosuke/pen/Pjoqqp) to convert a hex color code to a CSS filter:

For example, output for #00EE00 is

filter: invert(42%) sepia(93%) saturate(1352%) hue-rotate(87deg) brightness(119%) contrast(119%);

3. Add the CSS filter into this class.

   ```css
   .filter-green {
     filter: invert(48%) sepia(79%) saturate(2476%) hue-rotate(86deg) brightness(
         118%
       ) contrast(119%);
   }
   ```
