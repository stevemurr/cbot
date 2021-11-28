# CBOT 

CBOT is a coinbase trading framework and is a super work in progress.

- listens to coinbase ticker and saves price data
- **TODO**: try modeling trading [using ecs](https://www.youtube.com/watch?v=W3aieHjyNvw&list=PLSWK4JALZGZNVcTcoXcTjWn8DrUP7TOeR&index=2)

**Entity** is a collection of components  
**Component** is state with no behavior  
**System** has behavior and stores no state  

```
Entity Admin
    []System
    map[entityId]Entity
    []Component

EntityAdmin
    for system in systems
        system.update(timestep)

system.update
    for component in components
        component.update
```

# Run

pre-requisite: Install `make` and `golang` for your platform.

### Build
```
make build
```

### Run
```
make run
```

### Clean
```
make reset
```