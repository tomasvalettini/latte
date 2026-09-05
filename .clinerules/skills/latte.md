# Latte Skill

## Overview

Latte is a terminal-based drip tracker for managing task lists (called "blends") and individual tasks (called "drips"). It uses a simple CLI interface to list, add, update, and delete tasks.

## Naming Convention

- **Blends**: Task lists/groups (e.g., "Latte", "PokemonDB", "DIY")
- **Drips**: Individual tasks within a blend (indexed 0, 1, 2, etc.)

## Available Commands

### `latte list`

List all blends with their IDs and drip counts.

```bash
latte list
```

**Example output:**
```
===================================
            ALL BLENDS            
===================================
  [1] Programming Project (0 drips)
  [2] DIY (1 drip)
  [3] Latte (5 drips)
  [4] PokemonDB (3 drips)
```

### `latte list --blend-id <id>`

List all drips within a specific blend by its ID.

```bash
latte list --blend-id 3
```

**Example output:**
```
=====================
      Latte (3)      
=====================
  [0] move task saving to sqlite
  [1] create flutter app
  [2] create cli authentication
```

### `latte add "<drip description>" --blend-id <id>`

Add a new drip to a blend by its ID.

```bash
latte add "create project structure" --blend-id 3
```

### `latte add "<drip description>" --blend "<blend name>"`

Add a new drip to a blend by name. Creates the blend if it doesn't exist.

```bash
latte add "learn go" --blend "Latte"
```

### `latte delete --blend-id <id> --drip-id <drip-id>`

Delete a specific drip from a blend.

```bash
latte delete --blend-id 3 --drip-id 0
```

**Success message:**
```
Drip with id 0 deleted successfully.
```

### `latte update --blend-id <id> --drip-id <drip-id> "<new description>"`

Update a drip's description.

```bash
latte update --blend-id 3 --drip-id 1 "create authentication module"
```

## Common Patterns

### List all blends
```bash
latte list
```

### List drips in a specific blend
```bash
latte list --blend-id <blend_id>
```

### Add a drip to a blend
```bash
latte add "task description" --blend-id <blend_id>
```

### Delete a drip
```bash
latte delete --blend-id <blend_id> --drip-id <drip_id>
```

### Update a drip
```bash
latte update --blend-id <blend_id> --drip-id <drip_id> "new description"
```

## Flags Reference

| Flag | Description |
|------|-------------|
| `--blend-id int` | The ID of the blend (default -1) |
| `--blend string` | The blend name |
| `--drip-id int` | The ID of the drip to delete/update (default -1) |

## Quick Reference

```
latte list                          # All blends
latte list --blend-id 3             # Drips in blend 3
latte add "task" --blend-id 3       # Add task to blend 3
latte add "task" --blend "Latte"    # Add task to blend named "Latte"
latte delete --blend-id 3 --drip-id 0  # Delete drip 0 from blend 3
latte update --blend-id 3 --drip-id 0 "new task"  # Update drip
```