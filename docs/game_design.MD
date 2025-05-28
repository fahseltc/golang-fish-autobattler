# Game Design

## Fish (Item)

Each Item object sits in a container of Items called a Collection.
Each collection belongs to either a player or the current opponent.

Functionality:

- Activate(sourceItem, targetItem, data)
- TakeDamage(damage)
- IsAlive
- CollideWith

Data:

- Name
- Life
- CurrentLife
- Damage
- Duration
- CurrentDuration
- Sprite
