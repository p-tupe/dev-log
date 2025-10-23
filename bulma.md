---
modified: Fri Oct  7 13:07:29 EDT 2022
---
# Bulma

> Bulma is a free, open source framework that provides ready-to-use frontend components that you can easily combine to build responsive web interfaces.

- [Official Site](https://bulma.io/)

## Some useful classes

### Layout

- `container`: Centers content with some margin on the sides for bigger screens. Apply at top.

- `level`: A collection of vertically aligned items on a horizontal line.

  - `level is-mobile has-text-centered`: Similar to flex align and justify center on all screens
  - `level-item`: Level children class
  - `level-left`: To have children float left (on an intermediate element between `level` and `level-item`)
  - `level-right`: To have children float right (same as above)

- `hero`: Full-width banner

  - `hero is-fullheight`: Self-evident
  - `hero-body`: Is where inner content goes

- `section`: Logical parts as direct children of `body` elem

- `footer`: Container with padding at bottom

- `columns`: Other frameworks call it "rows". You get the picture.

  - `is-mobile`: to have columns spread out like desktop (instead of stacked on) on small devices
  - `is-vcentered`: to vertically align if non-uniform height in a row
  - `is-multiline`: wrap around below instead of shrinking or stretching
  - `is-centered`: Align all columns to center, useful with `is-mobile` and non-full-width columns
  - `column`: Children of rows, er, `columns`. Can be nested.
    - `is-<x>`: where `x` is a number between 1 to 12. Defines the width of column.
    - `is-half`: Half width column, same as `is-6`

### Elements

- `box`: A white container with some padding and shadow. Useful for holding stuff.

- `button`: Self-explainatory stuff

  - `is-[primary | secondary | info | success | warning | danger]`
  - `is-outlined`, `is-rounded`, `disabled`
  - `is-loading`: Adds a spinner to the button

- `image`: For responsive images

  - `is-rounded`
  - `is-<x>by<y>`: x and y define a ratio with values in 1,2,3,4,5,9,16.

- `table`, `tag`, `notification`, `delete`, `title | subtitle`, `progress` etc

### Components
