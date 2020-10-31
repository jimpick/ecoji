package ecoji

//The mappings for this file were generated in two steps.

/*
   #First, tan the following command to generate emojis.txt
   wget -qO- https://unicode.org/Public/emoji/11.0/emoji-test.txt | cut -f 1 -d ' ' | sort -u | sed '/^[#0]/ d' | sed '/^\s*$/d' > /tmp/emojis.txt
*/

/*

   //Second, used the following Java code to generate the mappings in this file
   try (Stream<String> stream = Files.lines(Paths.get("/tmp/emojis.txt"))) {
     List<String> emojis = stream.collect(Collectors.toList());

     System.out.println("const padding41 rune=0x" + emojis.remove(256));
     System.out.println("const padding42 rune=0x" + emojis.remove(512));
     System.out.println("const padding43 rune=0x" + emojis.remove(768));

     for (int i = 0; i < 1024; i++) {
       System.out.println("emojis[" + i + "]=0x" + emojis.get(i));
     }
   }
*/

//This should sort before everything.  This is output when 3 or less input bytes are present.
const padding rune = 0x2615

//The following paddings are used when only 4 of 5 input bytes are present.

//This should sort between padding and emojis[0]
const padding40 rune = 0x269C

//This should sort between emojis[255] and emojis[256]
const padding41 rune = 0x1F3CD

//This should sort between emojis[511] and emojis[512]
const padding42 rune = 0x1F4D1

//This should sort between emojis[767] and emojis[768]
const padding43 rune = 0x1F64B

const selectedGlyphs = "IJfijlrstºÌÍÎÏìíîïĨĩĪīĬĭĮįİıĴĵĺļľŀłŕŗřśŝşšţŧſƖƗƚƫƭǁǃǏǐǰȈȉȊȋȑȓșțɉḟṡṫẛʰʱʲʳʴʵʶʸʹʻʼʽʾʿˀˁ˂˃˄˅ˆˇˈˊˋˌˍˎˏːˑ˒˓˞˟ˠˡˢˣˤˬ˯˰˱˲˳˴˷˸˹˺˻˼˽˾ͱͿΐΙΪίιϊϳἰἱἲἳἴἵἶἷὶίῐῑῒΐῖῗῘῙІЇЈгзѓѕіїјґғҙӀӟꜣꝇꝉꝩꝲꝼꞁꞄꞅ꞉꞊Ꞌꞌɨɩɭɺɽɿʇʈᴉᴊᴬᴮᴯᴰᴱᴲᴳᴴᴵᴶᴷᴸᴹᴺᴻᴼᴽᴾᴿᵀᵁᵃᵄᵅᵇᵈᵉᵊᵋᵌᵍᵎᵏᵑᵒᵓᵔᵕᵖᵗᵘᵙᵛᵜᵝᵞᵟᵡᵢᵣᵤᵥᵦᵧᵨᵪᵮᵲᵳᵵᵻᵼﺍﻩﺁﺓءآ"

var emojis [1024]rune
var revEmojis map[rune]int

var narrowGlyphs [256]rune
var revNarrowGlyphs map[rune]int

func init() {

	emojis[0] = 0x1F004
	emojis[1] = 0x1F0CF
	emojis[2] = 0x1F170
	emojis[3] = 0x1F171
	emojis[4] = 0x1F17E
	emojis[5] = 0x1F17F
	emojis[6] = 0x1F18E
	emojis[7] = 0x1F191
	emojis[8] = 0x1F192
	emojis[9] = 0x1F193
	emojis[10] = 0x1F194
	emojis[11] = 0x1F195
	emojis[12] = 0x1F196
	emojis[13] = 0x1F197
	emojis[14] = 0x1F198
	emojis[15] = 0x1F199
	emojis[16] = 0x1F19A
	emojis[17] = 0x1F1E6
	emojis[18] = 0x1F1E7
	emojis[19] = 0x1F1E8
	emojis[20] = 0x1F1E9
	emojis[21] = 0x1F1EA
	emojis[22] = 0x1F1EB
	emojis[23] = 0x1F1EC
	emojis[24] = 0x1F1ED
	emojis[25] = 0x1F1EE
	emojis[26] = 0x1F1EF
	emojis[27] = 0x1F1F0
	emojis[28] = 0x1F1F1
	emojis[29] = 0x1F1F2
	emojis[30] = 0x1F1F3
	emojis[31] = 0x1F1F4
	emojis[32] = 0x1F1F5
	emojis[33] = 0x1F1F6
	emojis[34] = 0x1F1F7
	emojis[35] = 0x1F1F8
	emojis[36] = 0x1F1F9
	emojis[37] = 0x1F1FA
	emojis[38] = 0x1F1FB
	emojis[39] = 0x1F1FC
	emojis[40] = 0x1F1FD
	emojis[41] = 0x1F1FE
	emojis[42] = 0x1F1FF
	emojis[43] = 0x1F201
	emojis[44] = 0x1F202
	emojis[45] = 0x1F21A
	emojis[46] = 0x1F22F
	emojis[47] = 0x1F232
	emojis[48] = 0x1F233
	emojis[49] = 0x1F234
	emojis[50] = 0x1F235
	emojis[51] = 0x1F236
	emojis[52] = 0x1F237
	emojis[53] = 0x1F238
	emojis[54] = 0x1F239
	emojis[55] = 0x1F23A
	emojis[56] = 0x1F250
	emojis[57] = 0x1F251
	emojis[58] = 0x1F300
	emojis[59] = 0x1F301
	emojis[60] = 0x1F302
	emojis[61] = 0x1F303
	emojis[62] = 0x1F304
	emojis[63] = 0x1F305
	emojis[64] = 0x1F306
	emojis[65] = 0x1F307
	emojis[66] = 0x1F308
	emojis[67] = 0x1F309
	emojis[68] = 0x1F30A
	emojis[69] = 0x1F30B
	emojis[70] = 0x1F30C
	emojis[71] = 0x1F30D
	emojis[72] = 0x1F30E
	emojis[73] = 0x1F30F
	emojis[74] = 0x1F310
	emojis[75] = 0x1F311
	emojis[76] = 0x1F312
	emojis[77] = 0x1F313
	emojis[78] = 0x1F314
	emojis[79] = 0x1F315
	emojis[80] = 0x1F316
	emojis[81] = 0x1F317
	emojis[82] = 0x1F318
	emojis[83] = 0x1F319
	emojis[84] = 0x1F31A
	emojis[85] = 0x1F31B
	emojis[86] = 0x1F31C
	emojis[87] = 0x1F31D
	emojis[88] = 0x1F31E
	emojis[89] = 0x1F31F
	emojis[90] = 0x1F320
	emojis[91] = 0x1F321
	emojis[92] = 0x1F324
	emojis[93] = 0x1F325
	emojis[94] = 0x1F326
	emojis[95] = 0x1F327
	emojis[96] = 0x1F328
	emojis[97] = 0x1F329
	emojis[98] = 0x1F32A
	emojis[99] = 0x1F32B
	emojis[100] = 0x1F32C
	emojis[101] = 0x1F32D
	emojis[102] = 0x1F32E
	emojis[103] = 0x1F32F
	emojis[104] = 0x1F330
	emojis[105] = 0x1F331
	emojis[106] = 0x1F332
	emojis[107] = 0x1F333
	emojis[108] = 0x1F334
	emojis[109] = 0x1F335
	emojis[110] = 0x1F336
	emojis[111] = 0x1F337
	emojis[112] = 0x1F338
	emojis[113] = 0x1F339
	emojis[114] = 0x1F33A
	emojis[115] = 0x1F33B
	emojis[116] = 0x1F33C
	emojis[117] = 0x1F33D
	emojis[118] = 0x1F33E
	emojis[119] = 0x1F33F
	emojis[120] = 0x1F340
	emojis[121] = 0x1F341
	emojis[122] = 0x1F342
	emojis[123] = 0x1F343
	emojis[124] = 0x1F344
	emojis[125] = 0x1F345
	emojis[126] = 0x1F346
	emojis[127] = 0x1F347
	emojis[128] = 0x1F348
	emojis[129] = 0x1F349
	emojis[130] = 0x1F34A
	emojis[131] = 0x1F34B
	emojis[132] = 0x1F34C
	emojis[133] = 0x1F34D
	emojis[134] = 0x1F34E
	emojis[135] = 0x1F34F
	emojis[136] = 0x1F350
	emojis[137] = 0x1F351
	emojis[138] = 0x1F352
	emojis[139] = 0x1F353
	emojis[140] = 0x1F354
	emojis[141] = 0x1F355
	emojis[142] = 0x1F356
	emojis[143] = 0x1F357
	emojis[144] = 0x1F358
	emojis[145] = 0x1F359
	emojis[146] = 0x1F35A
	emojis[147] = 0x1F35B
	emojis[148] = 0x1F35C
	emojis[149] = 0x1F35D
	emojis[150] = 0x1F35E
	emojis[151] = 0x1F35F
	emojis[152] = 0x1F360
	emojis[153] = 0x1F361
	emojis[154] = 0x1F362
	emojis[155] = 0x1F363
	emojis[156] = 0x1F364
	emojis[157] = 0x1F365
	emojis[158] = 0x1F366
	emojis[159] = 0x1F367
	emojis[160] = 0x1F368
	emojis[161] = 0x1F369
	emojis[162] = 0x1F36A
	emojis[163] = 0x1F36B
	emojis[164] = 0x1F36C
	emojis[165] = 0x1F36D
	emojis[166] = 0x1F36E
	emojis[167] = 0x1F36F
	emojis[168] = 0x1F370
	emojis[169] = 0x1F371
	emojis[170] = 0x1F372
	emojis[171] = 0x1F373
	emojis[172] = 0x1F374
	emojis[173] = 0x1F375
	emojis[174] = 0x1F376
	emojis[175] = 0x1F377
	emojis[176] = 0x1F378
	emojis[177] = 0x1F379
	emojis[178] = 0x1F37A
	emojis[179] = 0x1F37B
	emojis[180] = 0x1F37C
	emojis[181] = 0x1F37D
	emojis[182] = 0x1F37E
	emojis[183] = 0x1F37F
	emojis[184] = 0x1F380
	emojis[185] = 0x1F381
	emojis[186] = 0x1F382
	emojis[187] = 0x1F383
	emojis[188] = 0x1F384
	emojis[189] = 0x1F385
	emojis[190] = 0x1F386
	emojis[191] = 0x1F387
	emojis[192] = 0x1F388
	emojis[193] = 0x1F389
	emojis[194] = 0x1F38A
	emojis[195] = 0x1F38B
	emojis[196] = 0x1F38C
	emojis[197] = 0x1F38D
	emojis[198] = 0x1F38E
	emojis[199] = 0x1F38F
	emojis[200] = 0x1F390
	emojis[201] = 0x1F391
	emojis[202] = 0x1F392
	emojis[203] = 0x1F393
	emojis[204] = 0x1F396
	emojis[205] = 0x1F397
	emojis[206] = 0x1F399
	emojis[207] = 0x1F39A
	emojis[208] = 0x1F39B
	emojis[209] = 0x1F39E
	emojis[210] = 0x1F39F
	emojis[211] = 0x1F3A0
	emojis[212] = 0x1F3A1
	emojis[213] = 0x1F3A2
	emojis[214] = 0x1F3A3
	emojis[215] = 0x1F3A4
	emojis[216] = 0x1F3A5
	emojis[217] = 0x1F3A6
	emojis[218] = 0x1F3A7
	emojis[219] = 0x1F3A8
	emojis[220] = 0x1F3A9
	emojis[221] = 0x1F3AA
	emojis[222] = 0x1F3AB
	emojis[223] = 0x1F3AC
	emojis[224] = 0x1F3AD
	emojis[225] = 0x1F3AE
	emojis[226] = 0x1F3AF
	emojis[227] = 0x1F3B0
	emojis[228] = 0x1F3B1
	emojis[229] = 0x1F3B2
	emojis[230] = 0x1F3B3
	emojis[231] = 0x1F3B4
	emojis[232] = 0x1F3B5
	emojis[233] = 0x1F3B6
	emojis[234] = 0x1F3B7
	emojis[235] = 0x1F3B8
	emojis[236] = 0x1F3B9
	emojis[237] = 0x1F3BA
	emojis[238] = 0x1F3BB
	emojis[239] = 0x1F3BC
	emojis[240] = 0x1F3BD
	emojis[241] = 0x1F3BE
	emojis[242] = 0x1F3BF
	emojis[243] = 0x1F3C0
	emojis[244] = 0x1F3C1
	emojis[245] = 0x1F3C2
	emojis[246] = 0x1F3C3
	emojis[247] = 0x1F3C4
	emojis[248] = 0x1F3C5
	emojis[249] = 0x1F3C6
	emojis[250] = 0x1F3C7
	emojis[251] = 0x1F3C8
	emojis[252] = 0x1F3C9
	emojis[253] = 0x1F3CA
	emojis[254] = 0x1F3CB
	emojis[255] = 0x1F3CC
	emojis[256] = 0x1F3CE
	emojis[257] = 0x1F3CF
	emojis[258] = 0x1F3D0
	emojis[259] = 0x1F3D1
	emojis[260] = 0x1F3D2
	emojis[261] = 0x1F3D3
	emojis[262] = 0x1F3D4
	emojis[263] = 0x1F3D5
	emojis[264] = 0x1F3D6
	emojis[265] = 0x1F3D7
	emojis[266] = 0x1F3D8
	emojis[267] = 0x1F3D9
	emojis[268] = 0x1F3DA
	emojis[269] = 0x1F3DB
	emojis[270] = 0x1F3DC
	emojis[271] = 0x1F3DD
	emojis[272] = 0x1F3DE
	emojis[273] = 0x1F3DF
	emojis[274] = 0x1F3E0
	emojis[275] = 0x1F3E1
	emojis[276] = 0x1F3E2
	emojis[277] = 0x1F3E3
	emojis[278] = 0x1F3E4
	emojis[279] = 0x1F3E5
	emojis[280] = 0x1F3E6
	emojis[281] = 0x1F3E7
	emojis[282] = 0x1F3E8
	emojis[283] = 0x1F3E9
	emojis[284] = 0x1F3EA
	emojis[285] = 0x1F3EB
	emojis[286] = 0x1F3EC
	emojis[287] = 0x1F3ED
	emojis[288] = 0x1F3EE
	emojis[289] = 0x1F3EF
	emojis[290] = 0x1F3F0
	emojis[291] = 0x1F3F3
	emojis[292] = 0x1F3F4
	emojis[293] = 0x1F3F5
	emojis[294] = 0x1F3F7
	emojis[295] = 0x1F3F8
	emojis[296] = 0x1F3F9
	emojis[297] = 0x1F3FA
	emojis[298] = 0x1F3FB
	emojis[299] = 0x1F3FC
	emojis[300] = 0x1F3FD
	emojis[301] = 0x1F3FE
	emojis[302] = 0x1F3FF
	emojis[303] = 0x1F400
	emojis[304] = 0x1F401
	emojis[305] = 0x1F402
	emojis[306] = 0x1F403
	emojis[307] = 0x1F404
	emojis[308] = 0x1F405
	emojis[309] = 0x1F406
	emojis[310] = 0x1F407
	emojis[311] = 0x1F408
	emojis[312] = 0x1F409
	emojis[313] = 0x1F40A
	emojis[314] = 0x1F40B
	emojis[315] = 0x1F40C
	emojis[316] = 0x1F40D
	emojis[317] = 0x1F40E
	emojis[318] = 0x1F40F
	emojis[319] = 0x1F410
	emojis[320] = 0x1F411
	emojis[321] = 0x1F412
	emojis[322] = 0x1F413
	emojis[323] = 0x1F414
	emojis[324] = 0x1F415
	emojis[325] = 0x1F416
	emojis[326] = 0x1F417
	emojis[327] = 0x1F418
	emojis[328] = 0x1F419
	emojis[329] = 0x1F41A
	emojis[330] = 0x1F41B
	emojis[331] = 0x1F41C
	emojis[332] = 0x1F41D
	emojis[333] = 0x1F41E
	emojis[334] = 0x1F41F
	emojis[335] = 0x1F420
	emojis[336] = 0x1F421
	emojis[337] = 0x1F422
	emojis[338] = 0x1F423
	emojis[339] = 0x1F424
	emojis[340] = 0x1F425
	emojis[341] = 0x1F426
	emojis[342] = 0x1F427
	emojis[343] = 0x1F428
	emojis[344] = 0x1F429
	emojis[345] = 0x1F42A
	emojis[346] = 0x1F42B
	emojis[347] = 0x1F42C
	emojis[348] = 0x1F42D
	emojis[349] = 0x1F42E
	emojis[350] = 0x1F42F
	emojis[351] = 0x1F430
	emojis[352] = 0x1F431
	emojis[353] = 0x1F432
	emojis[354] = 0x1F433
	emojis[355] = 0x1F434
	emojis[356] = 0x1F435
	emojis[357] = 0x1F436
	emojis[358] = 0x1F437
	emojis[359] = 0x1F438
	emojis[360] = 0x1F439
	emojis[361] = 0x1F43A
	emojis[362] = 0x1F43B
	emojis[363] = 0x1F43C
	emojis[364] = 0x1F43D
	emojis[365] = 0x1F43E
	emojis[366] = 0x1F43F
	emojis[367] = 0x1F440
	emojis[368] = 0x1F441
	emojis[369] = 0x1F442
	emojis[370] = 0x1F443
	emojis[371] = 0x1F444
	emojis[372] = 0x1F445
	emojis[373] = 0x1F446
	emojis[374] = 0x1F447
	emojis[375] = 0x1F448
	emojis[376] = 0x1F449
	emojis[377] = 0x1F44A
	emojis[378] = 0x1F44B
	emojis[379] = 0x1F44C
	emojis[380] = 0x1F44D
	emojis[381] = 0x1F44E
	emojis[382] = 0x1F44F
	emojis[383] = 0x1F450
	emojis[384] = 0x1F451
	emojis[385] = 0x1F452
	emojis[386] = 0x1F453
	emojis[387] = 0x1F454
	emojis[388] = 0x1F455
	emojis[389] = 0x1F456
	emojis[390] = 0x1F457
	emojis[391] = 0x1F458
	emojis[392] = 0x1F459
	emojis[393] = 0x1F45A
	emojis[394] = 0x1F45B
	emojis[395] = 0x1F45C
	emojis[396] = 0x1F45D
	emojis[397] = 0x1F45E
	emojis[398] = 0x1F45F
	emojis[399] = 0x1F460
	emojis[400] = 0x1F461
	emojis[401] = 0x1F462
	emojis[402] = 0x1F463
	emojis[403] = 0x1F464
	emojis[404] = 0x1F465
	emojis[405] = 0x1F466
	emojis[406] = 0x1F467
	emojis[407] = 0x1F468
	emojis[408] = 0x1F469
	emojis[409] = 0x1F46A
	emojis[410] = 0x1F46B
	emojis[411] = 0x1F46C
	emojis[412] = 0x1F46D
	emojis[413] = 0x1F46E
	emojis[414] = 0x1F46F
	emojis[415] = 0x1F470
	emojis[416] = 0x1F471
	emojis[417] = 0x1F472
	emojis[418] = 0x1F473
	emojis[419] = 0x1F474
	emojis[420] = 0x1F475
	emojis[421] = 0x1F476
	emojis[422] = 0x1F477
	emojis[423] = 0x1F478
	emojis[424] = 0x1F479
	emojis[425] = 0x1F47A
	emojis[426] = 0x1F47B
	emojis[427] = 0x1F47C
	emojis[428] = 0x1F47D
	emojis[429] = 0x1F47E
	emojis[430] = 0x1F47F
	emojis[431] = 0x1F480
	emojis[432] = 0x1F481
	emojis[433] = 0x1F482
	emojis[434] = 0x1F483
	emojis[435] = 0x1F484
	emojis[436] = 0x1F485
	emojis[437] = 0x1F486
	emojis[438] = 0x1F487
	emojis[439] = 0x1F488
	emojis[440] = 0x1F489
	emojis[441] = 0x1F48A
	emojis[442] = 0x1F48B
	emojis[443] = 0x1F48C
	emojis[444] = 0x1F48D
	emojis[445] = 0x1F48E
	emojis[446] = 0x1F48F
	emojis[447] = 0x1F490
	emojis[448] = 0x1F491
	emojis[449] = 0x1F492
	emojis[450] = 0x1F493
	emojis[451] = 0x1F494
	emojis[452] = 0x1F495
	emojis[453] = 0x1F496
	emojis[454] = 0x1F497
	emojis[455] = 0x1F498
	emojis[456] = 0x1F499
	emojis[457] = 0x1F49A
	emojis[458] = 0x1F49B
	emojis[459] = 0x1F49C
	emojis[460] = 0x1F49D
	emojis[461] = 0x1F49E
	emojis[462] = 0x1F49F
	emojis[463] = 0x1F4A0
	emojis[464] = 0x1F4A1
	emojis[465] = 0x1F4A2
	emojis[466] = 0x1F4A3
	emojis[467] = 0x1F4A4
	emojis[468] = 0x1F4A5
	emojis[469] = 0x1F4A6
	emojis[470] = 0x1F4A7
	emojis[471] = 0x1F4A8
	emojis[472] = 0x1F4A9
	emojis[473] = 0x1F4AA
	emojis[474] = 0x1F4AB
	emojis[475] = 0x1F4AC
	emojis[476] = 0x1F4AD
	emojis[477] = 0x1F4AE
	emojis[478] = 0x1F4AF
	emojis[479] = 0x1F4B0
	emojis[480] = 0x1F4B1
	emojis[481] = 0x1F4B2
	emojis[482] = 0x1F4B3
	emojis[483] = 0x1F4B4
	emojis[484] = 0x1F4B5
	emojis[485] = 0x1F4B6
	emojis[486] = 0x1F4B7
	emojis[487] = 0x1F4B8
	emojis[488] = 0x1F4B9
	emojis[489] = 0x1F4BA
	emojis[490] = 0x1F4BB
	emojis[491] = 0x1F4BC
	emojis[492] = 0x1F4BD
	emojis[493] = 0x1F4BE
	emojis[494] = 0x1F4BF
	emojis[495] = 0x1F4C0
	emojis[496] = 0x1F4C1
	emojis[497] = 0x1F4C2
	emojis[498] = 0x1F4C3
	emojis[499] = 0x1F4C4
	emojis[500] = 0x1F4C5
	emojis[501] = 0x1F4C6
	emojis[502] = 0x1F4C7
	emojis[503] = 0x1F4C8
	emojis[504] = 0x1F4C9
	emojis[505] = 0x1F4CA
	emojis[506] = 0x1F4CB
	emojis[507] = 0x1F4CC
	emojis[508] = 0x1F4CD
	emojis[509] = 0x1F4CE
	emojis[510] = 0x1F4CF
	emojis[511] = 0x1F4D0
	emojis[512] = 0x1F4D2
	emojis[513] = 0x1F4D3
	emojis[514] = 0x1F4D4
	emojis[515] = 0x1F4D5
	emojis[516] = 0x1F4D6
	emojis[517] = 0x1F4D7
	emojis[518] = 0x1F4D8
	emojis[519] = 0x1F4D9
	emojis[520] = 0x1F4DA
	emojis[521] = 0x1F4DB
	emojis[522] = 0x1F4DC
	emojis[523] = 0x1F4DD
	emojis[524] = 0x1F4DE
	emojis[525] = 0x1F4DF
	emojis[526] = 0x1F4E0
	emojis[527] = 0x1F4E1
	emojis[528] = 0x1F4E2
	emojis[529] = 0x1F4E3
	emojis[530] = 0x1F4E4
	emojis[531] = 0x1F4E5
	emojis[532] = 0x1F4E6
	emojis[533] = 0x1F4E7
	emojis[534] = 0x1F4E8
	emojis[535] = 0x1F4E9
	emojis[536] = 0x1F4EA
	emojis[537] = 0x1F4EB
	emojis[538] = 0x1F4EC
	emojis[539] = 0x1F4ED
	emojis[540] = 0x1F4EE
	emojis[541] = 0x1F4EF
	emojis[542] = 0x1F4F0
	emojis[543] = 0x1F4F1
	emojis[544] = 0x1F4F2
	emojis[545] = 0x1F4F3
	emojis[546] = 0x1F4F4
	emojis[547] = 0x1F4F5
	emojis[548] = 0x1F4F6
	emojis[549] = 0x1F4F7
	emojis[550] = 0x1F4F8
	emojis[551] = 0x1F4F9
	emojis[552] = 0x1F4FA
	emojis[553] = 0x1F4FB
	emojis[554] = 0x1F4FC
	emojis[555] = 0x1F4FD
	emojis[556] = 0x1F4FF
	emojis[557] = 0x1F500
	emojis[558] = 0x1F501
	emojis[559] = 0x1F502
	emojis[560] = 0x1F503
	emojis[561] = 0x1F504
	emojis[562] = 0x1F505
	emojis[563] = 0x1F506
	emojis[564] = 0x1F507
	emojis[565] = 0x1F508
	emojis[566] = 0x1F509
	emojis[567] = 0x1F50A
	emojis[568] = 0x1F50B
	emojis[569] = 0x1F50C
	emojis[570] = 0x1F50D
	emojis[571] = 0x1F50E
	emojis[572] = 0x1F50F
	emojis[573] = 0x1F510
	emojis[574] = 0x1F511
	emojis[575] = 0x1F512
	emojis[576] = 0x1F513
	emojis[577] = 0x1F514
	emojis[578] = 0x1F515
	emojis[579] = 0x1F516
	emojis[580] = 0x1F517
	emojis[581] = 0x1F518
	emojis[582] = 0x1F519
	emojis[583] = 0x1F51A
	emojis[584] = 0x1F51B
	emojis[585] = 0x1F51C
	emojis[586] = 0x1F51D
	emojis[587] = 0x1F51E
	emojis[588] = 0x1F51F
	emojis[589] = 0x1F520
	emojis[590] = 0x1F521
	emojis[591] = 0x1F522
	emojis[592] = 0x1F523
	emojis[593] = 0x1F524
	emojis[594] = 0x1F525
	emojis[595] = 0x1F526
	emojis[596] = 0x1F527
	emojis[597] = 0x1F528
	emojis[598] = 0x1F529
	emojis[599] = 0x1F52A
	emojis[600] = 0x1F52B
	emojis[601] = 0x1F52C
	emojis[602] = 0x1F52D
	emojis[603] = 0x1F52E
	emojis[604] = 0x1F52F
	emojis[605] = 0x1F530
	emojis[606] = 0x1F531
	emojis[607] = 0x1F532
	emojis[608] = 0x1F533
	emojis[609] = 0x1F534
	emojis[610] = 0x1F535
	emojis[611] = 0x1F536
	emojis[612] = 0x1F537
	emojis[613] = 0x1F538
	emojis[614] = 0x1F539
	emojis[615] = 0x1F53A
	emojis[616] = 0x1F53B
	emojis[617] = 0x1F53C
	emojis[618] = 0x1F53D
	emojis[619] = 0x1F549
	emojis[620] = 0x1F54A
	emojis[621] = 0x1F54B
	emojis[622] = 0x1F54C
	emojis[623] = 0x1F54D
	emojis[624] = 0x1F54E
	emojis[625] = 0x1F550
	emojis[626] = 0x1F551
	emojis[627] = 0x1F552
	emojis[628] = 0x1F553
	emojis[629] = 0x1F554
	emojis[630] = 0x1F555
	emojis[631] = 0x1F556
	emojis[632] = 0x1F557
	emojis[633] = 0x1F558
	emojis[634] = 0x1F559
	emojis[635] = 0x1F55A
	emojis[636] = 0x1F55B
	emojis[637] = 0x1F55C
	emojis[638] = 0x1F55D
	emojis[639] = 0x1F55E
	emojis[640] = 0x1F55F
	emojis[641] = 0x1F560
	emojis[642] = 0x1F561
	emojis[643] = 0x1F562
	emojis[644] = 0x1F563
	emojis[645] = 0x1F564
	emojis[646] = 0x1F565
	emojis[647] = 0x1F566
	emojis[648] = 0x1F567
	emojis[649] = 0x1F56F
	emojis[650] = 0x1F570
	emojis[651] = 0x1F573
	emojis[652] = 0x1F574
	emojis[653] = 0x1F575
	emojis[654] = 0x1F576
	emojis[655] = 0x1F577
	emojis[656] = 0x1F578
	emojis[657] = 0x1F579
	emojis[658] = 0x1F57A
	emojis[659] = 0x1F587
	emojis[660] = 0x1F58A
	emojis[661] = 0x1F58B
	emojis[662] = 0x1F58C
	emojis[663] = 0x1F58D
	emojis[664] = 0x1F590
	emojis[665] = 0x1F595
	emojis[666] = 0x1F596
	emojis[667] = 0x1F5A4
	emojis[668] = 0x1F5A5
	emojis[669] = 0x1F5A8
	emojis[670] = 0x1F5B1
	emojis[671] = 0x1F5B2
	emojis[672] = 0x1F5BC
	emojis[673] = 0x1F5C2
	emojis[674] = 0x1F5C3
	emojis[675] = 0x1F5C4
	emojis[676] = 0x1F5D1
	emojis[677] = 0x1F5D2
	emojis[678] = 0x1F5D3
	emojis[679] = 0x1F5DC
	emojis[680] = 0x1F5DD
	emojis[681] = 0x1F5DE
	emojis[682] = 0x1F5E1
	emojis[683] = 0x1F5E3
	emojis[684] = 0x1F5E8
	emojis[685] = 0x1F5EF
	emojis[686] = 0x1F5F3
	emojis[687] = 0x1F5FA
	emojis[688] = 0x1F5FB
	emojis[689] = 0x1F5FC
	emojis[690] = 0x1F5FD
	emojis[691] = 0x1F5FE
	emojis[692] = 0x1F5FF
	emojis[693] = 0x1F600
	emojis[694] = 0x1F601
	emojis[695] = 0x1F602
	emojis[696] = 0x1F603
	emojis[697] = 0x1F604
	emojis[698] = 0x1F605
	emojis[699] = 0x1F606
	emojis[700] = 0x1F607
	emojis[701] = 0x1F608
	emojis[702] = 0x1F609
	emojis[703] = 0x1F60A
	emojis[704] = 0x1F60B
	emojis[705] = 0x1F60C
	emojis[706] = 0x1F60D
	emojis[707] = 0x1F60E
	emojis[708] = 0x1F60F
	emojis[709] = 0x1F610
	emojis[710] = 0x1F611
	emojis[711] = 0x1F612
	emojis[712] = 0x1F613
	emojis[713] = 0x1F614
	emojis[714] = 0x1F615
	emojis[715] = 0x1F616
	emojis[716] = 0x1F617
	emojis[717] = 0x1F618
	emojis[718] = 0x1F619
	emojis[719] = 0x1F61A
	emojis[720] = 0x1F61B
	emojis[721] = 0x1F61C
	emojis[722] = 0x1F61D
	emojis[723] = 0x1F61E
	emojis[724] = 0x1F61F
	emojis[725] = 0x1F620
	emojis[726] = 0x1F621
	emojis[727] = 0x1F622
	emojis[728] = 0x1F623
	emojis[729] = 0x1F624
	emojis[730] = 0x1F625
	emojis[731] = 0x1F626
	emojis[732] = 0x1F627
	emojis[733] = 0x1F628
	emojis[734] = 0x1F629
	emojis[735] = 0x1F62A
	emojis[736] = 0x1F62B
	emojis[737] = 0x1F62C
	emojis[738] = 0x1F62D
	emojis[739] = 0x1F62E
	emojis[740] = 0x1F62F
	emojis[741] = 0x1F630
	emojis[742] = 0x1F631
	emojis[743] = 0x1F632
	emojis[744] = 0x1F633
	emojis[745] = 0x1F634
	emojis[746] = 0x1F635
	emojis[747] = 0x1F636
	emojis[748] = 0x1F637
	emojis[749] = 0x1F638
	emojis[750] = 0x1F639
	emojis[751] = 0x1F63A
	emojis[752] = 0x1F63B
	emojis[753] = 0x1F63C
	emojis[754] = 0x1F63D
	emojis[755] = 0x1F63E
	emojis[756] = 0x1F63F
	emojis[757] = 0x1F640
	emojis[758] = 0x1F641
	emojis[759] = 0x1F642
	emojis[760] = 0x1F643
	emojis[761] = 0x1F644
	emojis[762] = 0x1F645
	emojis[763] = 0x1F646
	emojis[764] = 0x1F647
	emojis[765] = 0x1F648
	emojis[766] = 0x1F649
	emojis[767] = 0x1F64A
	emojis[768] = 0x1F64C
	emojis[769] = 0x1F64D
	emojis[770] = 0x1F64E
	emojis[771] = 0x1F64F
	emojis[772] = 0x1F680
	emojis[773] = 0x1F681
	emojis[774] = 0x1F682
	emojis[775] = 0x1F683
	emojis[776] = 0x1F684
	emojis[777] = 0x1F685
	emojis[778] = 0x1F686
	emojis[779] = 0x1F687
	emojis[780] = 0x1F688
	emojis[781] = 0x1F689
	emojis[782] = 0x1F68A
	emojis[783] = 0x1F68B
	emojis[784] = 0x1F68C
	emojis[785] = 0x1F68D
	emojis[786] = 0x1F68E
	emojis[787] = 0x1F68F
	emojis[788] = 0x1F690
	emojis[789] = 0x1F691
	emojis[790] = 0x1F692
	emojis[791] = 0x1F693
	emojis[792] = 0x1F694
	emojis[793] = 0x1F695
	emojis[794] = 0x1F696
	emojis[795] = 0x1F697
	emojis[796] = 0x1F698
	emojis[797] = 0x1F699
	emojis[798] = 0x1F69A
	emojis[799] = 0x1F69B
	emojis[800] = 0x1F69C
	emojis[801] = 0x1F69D
	emojis[802] = 0x1F69E
	emojis[803] = 0x1F69F
	emojis[804] = 0x1F6A0
	emojis[805] = 0x1F6A1
	emojis[806] = 0x1F6A2
	emojis[807] = 0x1F6A3
	emojis[808] = 0x1F6A4
	emojis[809] = 0x1F6A5
	emojis[810] = 0x1F6A6
	emojis[811] = 0x1F6A7
	emojis[812] = 0x1F6A8
	emojis[813] = 0x1F6A9
	emojis[814] = 0x1F6AA
	emojis[815] = 0x1F6AB
	emojis[816] = 0x1F6AC
	emojis[817] = 0x1F6AD
	emojis[818] = 0x1F6AE
	emojis[819] = 0x1F6AF
	emojis[820] = 0x1F6B0
	emojis[821] = 0x1F6B1
	emojis[822] = 0x1F6B2
	emojis[823] = 0x1F6B3
	emojis[824] = 0x1F6B4
	emojis[825] = 0x1F6B5
	emojis[826] = 0x1F6B6
	emojis[827] = 0x1F6B7
	emojis[828] = 0x1F6B8
	emojis[829] = 0x1F6B9
	emojis[830] = 0x1F6BA
	emojis[831] = 0x1F6BB
	emojis[832] = 0x1F6BC
	emojis[833] = 0x1F6BD
	emojis[834] = 0x1F6BE
	emojis[835] = 0x1F6BF
	emojis[836] = 0x1F6C0
	emojis[837] = 0x1F6C1
	emojis[838] = 0x1F6C2
	emojis[839] = 0x1F6C3
	emojis[840] = 0x1F6C4
	emojis[841] = 0x1F6C5
	emojis[842] = 0x1F6CB
	emojis[843] = 0x1F6CC
	emojis[844] = 0x1F6CD
	emojis[845] = 0x1F6CE
	emojis[846] = 0x1F6CF
	emojis[847] = 0x1F6D0
	emojis[848] = 0x1F6D1
	emojis[849] = 0x1F6D2
	emojis[850] = 0x1F6E0
	emojis[851] = 0x1F6E1
	emojis[852] = 0x1F6E2
	emojis[853] = 0x1F6E3
	emojis[854] = 0x1F6E4
	emojis[855] = 0x1F6E5
	emojis[856] = 0x1F6E9
	emojis[857] = 0x1F6EB
	emojis[858] = 0x1F6EC
	emojis[859] = 0x1F6F0
	emojis[860] = 0x1F6F3
	emojis[861] = 0x1F6F4
	emojis[862] = 0x1F6F5
	emojis[863] = 0x1F6F6
	emojis[864] = 0x1F6F7
	emojis[865] = 0x1F6F8
	emojis[866] = 0x1F6F9
	emojis[867] = 0x1F910
	emojis[868] = 0x1F911
	emojis[869] = 0x1F912
	emojis[870] = 0x1F913
	emojis[871] = 0x1F914
	emojis[872] = 0x1F915
	emojis[873] = 0x1F916
	emojis[874] = 0x1F917
	emojis[875] = 0x1F918
	emojis[876] = 0x1F919
	emojis[877] = 0x1F91A
	emojis[878] = 0x1F91B
	emojis[879] = 0x1F91C
	emojis[880] = 0x1F91D
	emojis[881] = 0x1F91E
	emojis[882] = 0x1F91F
	emojis[883] = 0x1F920
	emojis[884] = 0x1F921
	emojis[885] = 0x1F922
	emojis[886] = 0x1F923
	emojis[887] = 0x1F924
	emojis[888] = 0x1F925
	emojis[889] = 0x1F926
	emojis[890] = 0x1F927
	emojis[891] = 0x1F928
	emojis[892] = 0x1F929
	emojis[893] = 0x1F92A
	emojis[894] = 0x1F92B
	emojis[895] = 0x1F92C
	emojis[896] = 0x1F92D
	emojis[897] = 0x1F92E
	emojis[898] = 0x1F92F
	emojis[899] = 0x1F930
	emojis[900] = 0x1F931
	emojis[901] = 0x1F932
	emojis[902] = 0x1F933
	emojis[903] = 0x1F934
	emojis[904] = 0x1F935
	emojis[905] = 0x1F936
	emojis[906] = 0x1F937
	emojis[907] = 0x1F938
	emojis[908] = 0x1F939
	emojis[909] = 0x1F93A
	emojis[910] = 0x1F93C
	emojis[911] = 0x1F93D
	emojis[912] = 0x1F93E
	emojis[913] = 0x1F940
	emojis[914] = 0x1F941
	emojis[915] = 0x1F942
	emojis[916] = 0x1F943
	emojis[917] = 0x1F944
	emojis[918] = 0x1F945
	emojis[919] = 0x1F947
	emojis[920] = 0x1F948
	emojis[921] = 0x1F949
	emojis[922] = 0x1F94A
	emojis[923] = 0x1F94B
	emojis[924] = 0x1F94C
	emojis[925] = 0x1F94D
	emojis[926] = 0x1F94E
	emojis[927] = 0x1F94F
	emojis[928] = 0x1F950
	emojis[929] = 0x1F951
	emojis[930] = 0x1F952
	emojis[931] = 0x1F953
	emojis[932] = 0x1F954
	emojis[933] = 0x1F955
	emojis[934] = 0x1F956
	emojis[935] = 0x1F957
	emojis[936] = 0x1F958
	emojis[937] = 0x1F959
	emojis[938] = 0x1F95A
	emojis[939] = 0x1F95B
	emojis[940] = 0x1F95C
	emojis[941] = 0x1F95D
	emojis[942] = 0x1F95E
	emojis[943] = 0x1F95F
	emojis[944] = 0x1F960
	emojis[945] = 0x1F961
	emojis[946] = 0x1F962
	emojis[947] = 0x1F963
	emojis[948] = 0x1F964
	emojis[949] = 0x1F965
	emojis[950] = 0x1F966
	emojis[951] = 0x1F967
	emojis[952] = 0x1F968
	emojis[953] = 0x1F969
	emojis[954] = 0x1F96A
	emojis[955] = 0x1F96B
	emojis[956] = 0x1F96C
	emojis[957] = 0x1F96D
	emojis[958] = 0x1F96E
	emojis[959] = 0x1F96F
	emojis[960] = 0x1F970
	emojis[961] = 0x1F973
	emojis[962] = 0x1F974
	emojis[963] = 0x1F975
	emojis[964] = 0x1F976
	emojis[965] = 0x1F97A
	emojis[966] = 0x1F97C
	emojis[967] = 0x1F97D
	emojis[968] = 0x1F97E
	emojis[969] = 0x1F97F
	emojis[970] = 0x1F980
	emojis[971] = 0x1F981
	emojis[972] = 0x1F982
	emojis[973] = 0x1F983
	emojis[974] = 0x1F984
	emojis[975] = 0x1F985
	emojis[976] = 0x1F986
	emojis[977] = 0x1F987
	emojis[978] = 0x1F988
	emojis[979] = 0x1F989
	emojis[980] = 0x1F98A
	emojis[981] = 0x1F98B
	emojis[982] = 0x1F98C
	emojis[983] = 0x1F98D
	emojis[984] = 0x1F98E
	emojis[985] = 0x1F98F
	emojis[986] = 0x1F990
	emojis[987] = 0x1F991
	emojis[988] = 0x1F992
	emojis[989] = 0x1F993
	emojis[990] = 0x1F994
	emojis[991] = 0x1F995
	emojis[992] = 0x1F996
	emojis[993] = 0x1F997
	emojis[994] = 0x1F998
	emojis[995] = 0x1F999
	emojis[996] = 0x1F99A
	emojis[997] = 0x1F99B
	emojis[998] = 0x1F99C
	emojis[999] = 0x1F99D
	emojis[1000] = 0x1F99E
	emojis[1001] = 0x1F99F
	emojis[1002] = 0x1F9A0
	emojis[1003] = 0x1F9A1
	emojis[1004] = 0x1F9A2
	emojis[1005] = 0x1F9B0
	emojis[1006] = 0x1F9B1
	emojis[1007] = 0x1F9B2
	emojis[1008] = 0x1F9B3
	emojis[1009] = 0x1F9B4
	emojis[1010] = 0x1F9B5
	emojis[1011] = 0x1F9B6
	emojis[1012] = 0x1F9B7
	emojis[1013] = 0x1F9B8
	emojis[1014] = 0x1F9B9
	emojis[1015] = 0x1F9C0
	emojis[1016] = 0x1F9C1
	emojis[1017] = 0x1F9C2
	emojis[1018] = 0x1F9D0
	emojis[1019] = 0x1F9D1
	emojis[1020] = 0x1F9D2
	emojis[1021] = 0x1F9D3
	emojis[1022] = 0x1F9D4
	emojis[1023] = 0x1F9D5

	revEmojis = make(map[rune]int)

	for i, r := range emojis {
		revEmojis[r] = i
	}

	var i int = 0
	revNarrowGlyphs = make(map[rune]int)
	for _, runeValue := range selectedGlyphs {
		narrowGlyphs[i] = runeValue
		revNarrowGlyphs[runeValue] = i
		// fmt.Printf("%v: %#U starts at byte position %d\n", i, runeValue, byteIndex)
		i++
	}

	/*
		const nihongo = "日本語"
		for index, runeValue := range nihongo {
			fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
		}
	*/
}
