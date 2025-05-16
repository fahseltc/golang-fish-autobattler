# Golang Fishgame (WIP title)

## Overall gameplay ideas

Something like Slay the Sprite/FTL/The Bazaar, but in the ocean.
Game starts with a selection of 3 different fish options to add to your 'school'.
From there, the player is offered an option of 3 choices repeatedly, with some sort of boss encounter every once in a while.
Combats are an option, the school will fight an array of other fish in an auto-battler type scene. When enemy fish are defeated, the might be able to be eaten by your fish for increased stats.

Fish can be rearranged in your school.
Fish can grow when fed a certain number of times and gain stats.
Some fish might eat other fish and replace them in your school.
Fish can have adjacency bonuses for being next to other fish.
Some fish might not be beneficial to its neighbors.

## Fish defensive mechanisms research

schooling - bonus damage (and life?) for adjacent schooling fish
camouflage - percent change to evade
speed + manueverability
spines - damage when attacked
poison - applies damage over time effect when attacking
mimicry - copy some other fish
flying - immune to damage for some period of time
safe burrows - digging / anenomes
electricity - stun when attacked

## Misc notes

different levels like Slay the spire

- The Shallows
- The Deep
- The Dark
- The Trench
  etc...

## Ocean Fish

Predators:

- Tuna - open ocean, migratory, very fast, good in the cold, cold blooded, must constantly swim, eats herring/mackerel, 150kg up to 680KG???!!!
- Cod - lives in areas with soft bottoms, eats small fish/lobsters, 50kg
- Salmon - rivers and coastal seas, eats small fish/invertibrates, 2-3 years of growth followed by a long migration to breed. Chinook salmon die after breeding, atlantic ones dont - 4kg usuall but up to 50kg seen
- Sturgeon - inland sea, very old, bony plates, valuable caviar - HUGE 1,571 kg
- Goby - rocky reefs, tiny predator, hide inside urchins, nearly always touch reef surface,pair mates,
- Anglerfish - deep sea, ambush predator,
  Marlin
  Swordfish
  Grunion - coastal open ocean, micro predator, reproduce on shore during full moon,

Filter Feeders:

- Mackerel - coastal to open ocean, filter feeder,

Pickers

- Clownfish - coral reef, plankton/algae picker, live in host anenome, changes sex throughout lifetime,

  Barracuda - open ocean and reef predator,
  Eel
  Puffer
  Seahorse - coral and rocky reefs, Visual planktivore (predator)
  snapper
  grouper
  Sunfish - open ocean, foraging predator
  herring
  sardine
  lionfish
  flounder
  dolphin
  whale
  manta ray

## Eventing

Eventing might be useful in this project, look into this library:
https://github.com/maniartech/signals
