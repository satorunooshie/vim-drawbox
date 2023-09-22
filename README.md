# vim-drawbox - Draw ASCII diagrams in Vim
[gyim/vim-boxdraw](https://github.com/gyim/vim-boxdraw)'s Go version.

[![See vim-boxdraw in action](https://asciinema.org/a/qeig6TH6N4uteq7J6n4epUGaq.png)](https://asciinema.org/a/qeig6TH6N4uteq7J6n4epUGaq)

vim-drawbox plugin enable to draw ASCII diagrams in `blockwise-visual` mode.

## Usage
1. Select a rectangle via blockwise-visual mode.
2. Invoke a draw command.

## Installation
This plugin executes [drawbox](https://github.com/satorunooshie/vim-drawbox/tree/main/cmd).

The binary can be installed from [Releases](https://github.com/satorunooshie/vim-drawbox/releases) or use the following command if already installed [Go](https://github.com/golang/go).

```
cd cmd && go install
```

## Command List

| Command               | Description                                       |
|-----------------------|---------------------------------------------------|
| `+o`                  | Draw a rectangle.                                 |
| `+O` or `+mcb`        | Draw a labeled rectangle.                         |
| `+[O` or `+mlb`       | Draw a labeled rectangle with top-left label.     |
| `+]O` or `+mrb`       | Draw a labeled rectangle with top-right label.    |
| `+{O` or `+tcb`       | Draw a labeled rectangle with top-center label.   |
| `+}O` or `+bcb`       | Draw a labeled rectangle with bottom-center label.|
| `+{[O` or `+tlb`      | Draw a labeled rectangle with top-left label.     |
| `+{]O` or `+trb`      | Draw a labeled rectangle with top-right label.    |
| `+}[O` or `+blb`      | Draw a labeled rectangle with bottom-left label.  |
| `+}]O` or `+brb`      | Draw a labeled rectangle with bottom-right label. |
| `+c` or `+mcl`        | Place a label on the middle center.               |
| `+[c` or `+mll`       | Place a label on the middle left.                 |
| `+]c` or `+mrl`       | Place a label on the middle right.                |
| `+{c` or `+tcl`       | Place a label on the top center.                  |
| `+}c` or `+bcl`       | Place a label on the bottom center.               |
| `+{[c` or `+tll`      | Place a label on the top left.                    |
| `+{]c` or `+trl`      | Place a label on the top right.                   |
| `+}[c` or `+bll`      | Place a label on the bottom left.                 |
| `+}]c` or `+brl`      | Place a label on the bottom right.                |
| `+>` or `+<`          | Draw a horizontal line.                           |
| `+V`, `+v`, or `+^`   | Draw a vertical line.                             |
| `++>` or `++<`        | Draw a double horizontal line.                    |
| `++V`, `++v`, or `++^`| Draw a double vertical line.                      |
| `+-` or `+_`          | Draw a horizontal line.                           |
| `+\|`                  | Draw a vertical line.                             |
| `ao` or `ab`          | Select the outer box.                             |
| `io` or `ib`          | Select the inner box.                             |

## License
MIT
