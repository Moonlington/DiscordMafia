package main

import (
	"math/rand"
	"strings"
)

// These generators are not mine, they are from fantasynamegenerators.com
// I remade some in golang for the bot.

func cityName() string {
	var names1 = []string{"a", "e", "i", "o", "u", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	var names2 = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z", "br", "cr", "dr", "fr", "gr", "kr", "pr", "qr", "sr", "tr", "vr", "wr", "yr", "zr", "str", "bl", "cl", "fl", "gl", "kl", "pl", "sl", "vl", "yl", "zl", "ch", "kh", "ph", "sh", "yh", "zh"}
	var names3 = []string{"a", "e", "i", "o", "u", "ae", "ai", "ao", "au", "aa", "ee", "ea", "ei", "eo", "eu", "ia", "ie", "io", "iu", "oa", "oe", "oi", "oo", "ou", "ua", "ue", "ui", "uo", "uu", "a", "e", "i", "o", "u", "a", "e", "i", "o", "u", "a", "e", "i", "o", "u", "a", "e", "i", "o", "u", "a", "e", "i", "o", "u"}
	var names4 = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z", "br", "cr", "dr", "fr", "gr", "kr", "pr", "tr", "vr", "wr", "zr", "st", "bl", "cl", "fl", "gl", "kl", "pl", "sl", "vl", "zl", "ch", "kh", "ph", "sh", "zh"}
	var names5 = []string{"c", "d", "f", "h", "k", "l", "m", "n", "p", "r", "s", "t", "x", "y", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	var names6 = []string{"aco", "ada", "adena", "ago", "agos", "aka", "ale", "alo", "am", "anbu", "ance", "and", "ando", "ane", "ans", "anta", "arc", "ard", "ares", "ario", "ark", "aso", "athe", "eah", "edo", "ego", "eigh", "eim", "eka", "eles", "eley", "ence", "ens", "ento", "erton", "ery", "esa", "ester", "ey", "ia", "ico", "ido", "ila", "ille", "in", "inas", "ine", "ing", "irie", "ison", "ita", "ock", "odon", "oit", "ok", "olis", "olk", "oln", "ona", "oni", "onio", "ont", "ora", "ord", "ore", "oria", "ork", "osa", "ose", "ouis", "ouver", "ul", "urg", "urgh", "ury"}
	var names7 = []string{"bert", "bridge", "burg", "burgh", "burn", "bury", "bus", "by", "caster", "cester", "chester", "dale", "dence", "diff", "ding", "don", "fast", "field", "ford", "gan", "gas", "gate", "gend", "ginia", "gow", "ham", "hull", "land", "las", "ledo", "lens", "ling", "mery", "mond", "mont", "more", "mouth", "nard", "phia", "phis", "polis", "pool", "port", "pus", "ridge", "rith", "ron", "rora", "ross", "rough", "sa", "sall", "sas", "sea", "set", "sey", "shire", "son", "stead", "stin", "ta", "tin", "tol", "ton", "vale", "ver", "ville", "vine", "ving", "well", "wood"}
	var name string
	i := rand.Intn(11)
	if i < 3 {
		rnd0 := int(rand.Float64() * float64(len(names1)))
		rnd1 := int(rand.Float64() * float64(len(names2)))
		rnd2 := int(rand.Float64() * float64(len(names3)))
		rnd3 := int(rand.Float64() * float64(len(names5)))
		rnd4 := int(rand.Float64() * float64(len(names7)))
		name = names1[rnd0] + names2[rnd1] + names3[rnd2] + names5[rnd3] + names7[rnd4]
	} else if i < 5 {
		rnd2 := int(rand.Float64() * float64(len(names3)))
		rnd3 := int(rand.Float64() * float64(len(names4)))
		rnd4 := int(rand.Float64() * float64(len(names3)))
		rnd5 := int(rand.Float64() * float64(len(names5)))
		rnd6 := int(rand.Float64() * float64(len(names7)))
		name = names3[rnd2] + names4[rnd3] + names3[rnd4] + names5[rnd5] + names7[rnd6]
	} else if i < 8 {
		rnd0 := int(rand.Float64() * float64(len(names1)))
		rnd1 := int(rand.Float64() * float64(len(names2)))
		rnd2 := int(rand.Float64() * float64(len(names6)))
		name = names1[rnd0] + names2[rnd1] + names6[rnd2]
	} else {
		rnd0 := int(rand.Float64() * float64(len(names1)))
		rnd1 := int(rand.Float64() * float64(len(names2)))
		rnd2 := int(rand.Float64() * float64(len(names3)))
		rnd3 := int(rand.Float64() * float64(len(names4)))
		rnd4 := int(rand.Float64() * float64(len(names6)))
		name = names1[rnd0] + names2[rnd1] + names3[rnd2] + names4[rnd3] + names6[rnd4]
	}
	return strings.Title(name)
}

func companyName() string {
	var names1 = []string{"Accent", "Ace", "Alligator", "Alp", "Alpha", "Alpine", "Amazon", "Angel", "Answer", "Ant", "Antler", "Apache", "Apex", "Apple", "Apricot", "Arcane", "Ask", "Aura", "Banshee", "Bear Paw", "Beedle", "Berry", "Beta", "Blizzard", "Blossom", "Blue", "Bluff", "Boar", "Bridge", "Brisk", "Buck", "Bull", "Bumblebee", "Butterfly", "Buzzy Bee", "Canine", "Cannon", "Cave", "Cavern", "Ceasar", "Champion", "Chief", "Cliff", "Cloud", "Clover", "Core", "Crow", "Cruise", "Crux", "Cryptic", "Crystal", "Cube", "Cyclone", "Cyclops", "Daydream", "Deluge", "Desert", "Diamond", "Dinosaur", "Dragon", "Dream", "Drift", "Dwarf", "Dynamic", "Eclipse", "Ecstatic", "Electra", "Electron", "Elite", "Elvish", "Energy", "Enigma", "Equinox", "Essence", "Explorer", "Fairy", "Feline", "Fire", "Fjord", "Flower", "Fluke", "Flux", "Forest", "Fortune", "Freak", "Frostfire", "Gale", "Gem", "Ghost", "Glacier", "Global", "Globe", "Gnome", "Goblin", "Gold", "Golden Road", "Gorilla", "Granite", "Grasshopper", "Great White", "Green", "Griffin", "Grizzly", "Grotto", "Hammer", "Happy", "Hatch", "Heart", "Herb", "Hercules", "Hero", "High Tide", "Hog", "Honey", "Honeydew", "Hook", "Hound", "Hummingbird", "Hurricane", "Ice", "Iceberg", "Icecap", "Imagination", "Imagine", "Iron", "Jet", "Jewel", "Joy", "Jungle", "Jupiter", "Karma", "Labyrinth", "Lagoon", "Lemon", "Leopard", "Life", "Lion", "Lioness", "Loki", "Low Tide", "Lucent", "Lucky", "Maple", "Marble", "Mars", "Marsh", "Maze", "Melon", "Mercury", "Mermaid", "Micro", "Midnight", "Monarch", "Moon", "Moondust", "Moonlight", "Moonlit", "Moth", "Motion", "Mount", "Mountain", "Mystic", "Nemo", "Neptune", "Nero", "Night", "Nimble", "North Star", "Nymph", "Oak", "Ocean", "Odin", "Ogre", "Omega", "Oracle", "Orange", "Orc", "Owl", "Oyster", "Padlock", "Parable", "Paragon", "Peach", "Pearl", "Petal", "Phantasm", "Phantom", "Phenomenon", "Phoenix", "Pink", "Pinnacle", "Pixel", "Pixy", "Pluto", "Pride", "Primal", "Prime", "Prodigy", "Prophecy", "Proton", "Pumpkin", "Purple", "Pyramid", "Quad", "Question", "Rabbit", "Radiant", "Raptor", "Raven", "Red", "Revelation", "Rhino", "Riddle", "Ridge", "River", "Robin", "Root", "Rose", "Rush", "Sail", "Sapling", "Saturn", "Saw", "Seaway", "Seed", "Shade", "Shadow", "Sharkfin", "Shrub", "Sign", "Signal", "Silver", "Silver Lining", "Slick", "Smart", "Smile", "Solstice", "Soul", "Specter", "Sphere", "Sphinx", "Spider", "Spike", "Spirit", "Sprite", "Squid", "Star", "Stardust", "Starlight", "Storm", "Summit", "Sun", "Sunlight", "Sunshine", "Surge", "Surprise", "Tempest", "Thor", "Thunder", "Tide", "Tiger", "Tigress", "Timber", "Titanium", "Topiary", "Tortoise", "Trek", "Triumph", "Tucan", "Tulip", "Tundra", "Turtle", "Twilight", "Twister", "Typhoon", "Valkyrie", "Venus", "Vertex", "Vine", "Vision", "Void", "Vortex", "Voyage", "Wave", "Web", "Whirlpool", "Whirlwind", "White Wolf", "Whiteout", "Whiz", "Willow", "Witch", "Wizard", "Wolf", "Wonder", "Wood", "World", "Yellow", "Yew", "Zeus"}
	var names2 = []string{"", "", "Acoustics", "Arts", "Aviation", "Brews", "Co.", "Coms", "Corp", "Corporation", "Electronics", "Enterprises", "Entertainment", "Foods", "Industries", "Intelligence", "Lighting", "Limited", "Media", "Microsystems", "Motors", "Navigations", "Networks", "Productions", "Records", "Security", "Softwares", "Solutions", "Sports", "Systems", "Technologies"}
	var names3 = []string{"Accenco", "Aces", "Allico", "Alphacom", "Alpire", "Alpite", "Amazystems", "Angelico", "Ansoft", "Antarts", "Antelligence", "Apachicorp", "Apexi", "Appiation", "Aprico", "Arcanetworks", "Asco", "Aurarts", "Bansheelectronis", "Bearings", "Beedlectrics", "Berrycords", "Betarts", "Blizzart", "Blossomotors", "Bluetronics", "Bluffoods", "Boarco", "Bridgelectrics", "Brisco", "Buckoustics", "Bullimited", "Bumblebeelectrics", "Butterflyght", "Buzzylectrics", "Canics", "Cannonics", "Cavernetworks", "Cavologies", "Ceasarts", "Champroductions", "Chiefoods", "Cliffoods", "Cloudustries", "Cloverprises", "Corecords", "Crowares", "Cruisertainment", "Cruxolutions", "Crypticorps", "Crystalightning", "Cubrews", "Cyclolutions", "Cycloration", "Daydreamotors", "Delugation", "Deserprises", "Diamontronics", "Dinosecurity", "Dragonetworks", "Dreamedia", "Driftonics", "Dwarfoods", "Dynamico", "Ecliprises", "Ecstaticorps", "Elecoms", "Electrorks", "Elitelligence", "Elviations", "Energence", "Enigmotors", "Equinetworks", "Essensecurity", "Explority", "Fairiprises", "Felinetworks", "Firetronics", "Fjordustries", "Flowertainment", "Flukords", "Fluxystems", "Forestics", "Fortunetworks", "Freacrosystems", "Frostfiretronics", "Galerprises", "Gemedia", "Ghostronics", "Glaciarts", "Globaviations", "Globeworks", "Gnomelectrics", "Goblintelligence", "Golbrews", "Goldustries", "Gorillacoustics", "Granitelligence", "Grasshoproductions", "Greatechnologies", "Greenetworks", "Griffindustries", "Grizzlimited", "Grottolutions", "Hammertronics", "Happindustries", "Hatchworks", "Heartelligence", "Herbrews", "Herculentertainment", "Herolutions", "High Tidustries", "Hogurity", "Honeydustries", "Honeytelligence", "Hookurity", "Houndnavigations", "Hummingbirdustries", "Hurricanetworks", "Icebergarts", "Icecaproductions", "Icecorps", "Imaginavigations", "Imaginetworks", "Ironavigation", "Jetechnologies", "Jewelectrics", "Joytechs", "Junglectics", "Jupitelligence", "Karmarts", "Labyrintelligence", "Lagoonavigations", "Lemonetworks", "Leopardworks", "Lifoods", "Lionessolutions", "Lionetworks", "Lokilutions", "Low Tidustries", "Lucentertainment", "Luckytronics", "Mapletainment", "Marblightning", "Marsecuriy", "Marsoftwares", "Mazecurity", "Melonetworks", "Mercurtainment", "Mermedia", "Microlutions", "Midnightelligence", "Monarctronics", "Moondustries", "Moonlightings", "Moonlimited", "Moonnetworks", "Mothechnologies", "Motionavigations", "Mountainetworks", "Mountelligence", "Mysticorps", "Nemotors", "Neptunetworks", "Neroductions", "Nightelligence", "Nimbletainment", "North Starporation", "Nymphoods", "Oakoms", "Oceanavigations", "Odinetworks", "Ogreprises", "Omegacoustics", "Oracleutions", "Orangations", "Orco", "Owlimited", "Oystertainment", "Padlockurity", "Parableutions", "Paragonetworks", "Peachco", "Pearlightning", "Petalimited", "Phantasmedia", "Phantomedia", "Phenomenologies", "Phoenixolutions", "Pinkorps", "Pinnaclelectrics", "Pixelimited", "Pixystems", "Plutronics", "Priductions", "Primacoustics", "Primedia", "Prodintelligence", "Prophecycurity", "Protonetworks", "Pumpkinavigation", "Purplelimited", "Pyramidustries", "Quaductions", "Questindustries", "Rabbitechnologies", "Radiantelligence", "Raptolutions", "Ravenetworks", "Redustries", "Revelationetworks", "Rhinotainment", "Riddlectronics", "Ridgeco", "Riverecords", "Robinetworks", "Rootechnologies", "Rosecurity", "Rushcorp", "Sailightning", "Saplimited", "Saturnetworks", "Sawwares", "Seawares", "Seedtronics", "Shadearts", "Shadoworks", "Sharkfinetworks", "Shrubrews", "Signalimited", "Signetworks", "Silver Linetworks", "Silverecords", "Slickorps", "Smartechnologies", "Smilectronics", "Solsticetems", "Soulightning", "Spectertainment", "Spherecords", "Sphinxecurity", "Spicurity", "Spideradio", "Spiritechnologies", "Spritechnologies", "Squidustries", "Stardustechnologies", "Starecords", "Starlightning", "Stormedia", "Summitechnologies", "Sunavigations", "Sunlightning", "Sunshinetworks", "Surgesystems", "Surprisystems", "Tempestechnologies", "Thorecords", "Thunderecords", "Tidustries", "Tigertainment", "Tigressystems", "Timberecords", "Titaniumotors", "Topiarytelligence", "Tortoisecurity", "Trekords", "Triumphoods", "Tucanterprises", "Tuliproductions", "Tundracoustics", "Turtletainment", "Twilightechnologies", "Twisterecords", "Typhoonavigations", "Valkyrecords", "Venusystems", "Vertexoftwards", "Vinedustries", "Visionetworks", "Voidustries", "Vortexecurity", "Voyagetronics", "Wavigations", "Webrews", "Whirlpoolutions", "Whirlwindustries", "White Wolfoods", "Whiteoutwares", "Whizystems", "Willowares", "Witchystems", "Wizardustries", "Wolfoods", "Wonderprises", "Wooductions", "Worldwares", "Yelloworks", "Yeworks", "Zeusolutions"}
	var names4 = []string{"Accent", "Ace", "Alligator", "Alpha", "Alpine", "Amazon", "Angel", "Apache", "Apex", "Arcane", "Aura", "Banshee", "Beedle", "Berry", "Beta", "Blossom", "Blue", "Boar", "Bridge", "Bull", "Bee", "Cannon", "Cave", "Cavern", "Chief", "Cliff", "Cloud", "Core", "Crow", "Crystal", "Cube", "Desert", "Diamond", "Dino", "Dragon", "Dream", "Drift", "Dwarf", "Dynamic", "Fairy", "Fire", "Flower", "Forest", "Freak", "Gem", "Ghost", "Global", "Globe", "Gnome", "Gold", "Gorilla", "Granite", "Green", "Griffin", "Grizzly", "Grotto", "Hammer", "Happy", "Hatch", "Heart", "Herb", "Hero", "Hog", "Honey", "Hook", "Hound", "Humming", "Ice", "Iron", "Jet", "Joy", "Karma", "Lagoon", "Lemon", "Leopard", "Life", "Lion", "Lioness", "Lucent", "Lucky", "Maple", "Marble", "Mars", "Marsh", "Maze", "Mermaid", "Micro", "Moon", "Motion", "Mountain", "Nemo", "Night", "Nimble", "Nymph", "Oak", "Ocean", "Omega", "Owl", "Oyster", "Peach", "Pearl", "Petal", "Phantomn", "Phoenix", "Pink", "Pinnacle", "Pixel", "Pixy", "Prime", "Purple", "Quad", "Raven", "Red", "Rhino", "Riddle", "Ridge", "River", "Robin", "Root", "Rose", "Sail", "Shade", "Shadow", "Sign", "Signal", "Silver", "Smart", "Smile", "Soul", "Sphere", "Spider", "Spike", "Spirit", "Sprite", "Squid", "Star", "Storm", "Sun", "Thunder", "Tiger", "Timber", "Tulip", "Twilight", "Typhoon", "Vine", "Void", "Vortex", "Wave", "Web", "Wizard", "Wolf", "Wonder", "Wood", "World", "Yellow", "Yew", "Zeus"}
	var names5 = []string{"aid", "air", "bank", "bar", "beat", "bit", "bite", "books", "bridge", "cast", "cloud", "coms", "dale", "dew", "dream", "ex", "fly", "force", "fruit", "gate", "gold", "head", "hive", "house", "hut", "king", "land", "life", "light", "man", "mart", "master", "media", "mobile", "nite", "paw", "phone", "point", "poly", "rover", "scape", "search", "shack", "shade", "shadow", "shine", "show", "space", "star", "stone", "stones", "sun", "sys", "tales", "techs", "time", "tronics", "tube", "walk", "ware", "wares", "water", "way", "ways", "well", "wheels", "wood", "works", "world", "worth"}
	var name string
	i := rand.Intn(11)
	if i < 3 {
		rnd0 := int(rand.Float64() * float64(len(names1)))
		rnd1 := int(rand.Float64() * float64(len(names2)))
		name = names1[rnd0] + " " + names2[rnd1]
	} else if i < 7 {
		rnd0 := int(rand.Float64() * float64(len(names3)))
		name = names3[rnd0]
	} else {
		rnd0 := int(rand.Float64() * float64(len(names4)))
		rnd1 := int(rand.Float64() * float64(len(names5)))
		name = names4[rnd0] + names5[rnd1]
	}
	return name
}
