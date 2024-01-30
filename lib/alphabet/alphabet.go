package alphabet

// Vowels

// অ BENGALI LETTER A
const V_A = 0x0985

// আ BENGALI LETTER AA
const V_AA = 0x0986

// ই BENGALI LETTER I
const V_I = 0x0987

// ঈ BENGALI LETTER II
const V_II = 0x0988

// উ BENGALI LETTER U
const V_U = 0x0989

// ঊ BENGALI LETTER UU
const V_UU = 0x098A

// ঋ BENGALI LETTER VOCALIC R
const V_R = 0x098B

// ঌ BENGALI LETTER VOCALIC L
// * Deprecated in Modern Bengali
const V_L = 0x098C


// এ BENGALI LETTER E
const V_E = 0x098F

// ঐ BENGALI LETTER AI
const V_EE = 0x0990

// ও BENGALI LETTER O
const V_O = 0x0993

// ঔ BENGALI LETTER AU
const V_OU = 0x0994

var VOWELS_ALL []uint16 = []uint16{
	V_A,
	V_AA,
	V_I,
	V_II,
	V_U,
	V_UU,
	V_R,
	V_E,
	V_EE,
	V_O,
	V_OU,
}

// Consonants

// ক BENGALI LETTER KA
const C_K = 0x0995

// খ BENGALI LETTER KHA
const C_KH = 0x0996

// গ BENGALI LETTER GA
const C_G = 0x0997

// ঘ BENGALI LETTER GHA
const C_GH = 0x0998

// ঙ BENGALI LETTER NGA
const C_U = 0x0999

// চ BENGALI LETTER CA
const C_C = 0x099A

// ছ BENGALI LETTER CHA
const C_CH = 0x099B

// জ BENGALI LETTER JA
const C_J = 0x099C

// ঝ BENGALI LETTER JHA
const C_JH = 0x099D

// ঞ BENGALI LETTER NYA
const C_NY = 0x099E

// ট BENGALI LETTER TTA
const C_T = 0x099F

// ঠ BENGALI LETTER TTHA
const C_TH = 0x09A0

// ড BENGALI LETTER DDA
const C_D = 0x09A1

// ঢ BENGALI LETTER DDHA
const C_DH = 0x09a2

// ণ BENGALI LETTER NNA
const C_NN = 0x09a3

// ত BENGALI LETTER TA
const C_TA = 0x09a4

// থ BENGALI LETTER THA
const C_THA = 0x09a5

// দ BENGALI LETTER DA
const C_DA = 0x09a6

// ধ BENGALI LETTER DHA
const C_DHA = 0x09a7

// ন BENGALI LETTER NA
const C_N = 0x09a8

// প BENGALI LETTER PA
const C_P = 0x09aa

// ফ BENGALI LETTER PHA
const C_PH = 0x09ab

// ব BENGALI LETTER BA
const C_B = 0x09ac

// ভ BENGALI LETTER BHA
const C_BH = 0x09ad

// ম BENGALI LETTER MA
const C_M = 0x09ae

// য BENGALI LETTER YA
const C_Z = 0x09af

// র BENGALI LETTER RA
const C_R = 0x09b0

// ল BENGALI LETTER LA
const C_L = 0x09b2

// শ BENGALI LETTER SHA
const C_SH = 0x09b6

// ষ BENGALI LETTER SSA
const C_SSH = 0x09b7

// স BENGALI LETTER SA
const C_S = 0x0b8

// হ BENGALI LETTER HA
const C_H = 0x09b9

// ড় BENGALI LETTER RRA
const C_RH = 0x09dc

// ঢ় BENGALI LETTER RHA
const C_RHH = 0x09dd

// য় BENGALI LETTER YYA
const C_Y = 0x09df

// ৎ BENGALI LETTER KHANDA TA
const C_KT = 0x09ce

// $ং BENGALI SIGN ANUSVARA
const C_NG = 0x0982

// $ঃ BENGALI SIGN VISARGA
const C_HH = 0x0983

// $ঁ BENGALI SIGN CANDRABINDU
const C_NNG = 0x0981

var CONSONANTS_ALL []uint16 = []uint16{
	C_K,
	C_KH,
	C_G,
	C_GH,
	C_U,
	C_C,
	C_CH,
	C_J,
	C_JH,
	C_NY,
	C_T,
	C_TH,
	C_D,
	C_DH,
	C_NN,
	C_TA,
	C_THA,
	C_DA,
	C_DHA,
	C_N,
	C_P,
	C_PH,
	C_B,
	C_BH,
	C_M,
	C_Z,
	C_R,
	C_L,
	C_SH,
	C_SSH,
	C_S,
	C_H,
	C_RH,
	C_RHH,
	C_Y,
	C_KT,
	C_NG,
	C_HH,
	C_NNG,
}
