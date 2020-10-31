console.log('Hi Jim!')

const glyphSelectors = [
  'table.wikitable:nth-of-type(3) tr td:nth-of-type(2)', // Basic Latin
  'table.wikitable:nth-of-type(4) tr td:nth-of-type(2)', // Latin-1 Supplement
  'table.wikitable:nth-of-type(5) tr td:nth-of-type(2)', // Latin Extended-A
  'table.wikitable:nth-of-type(6) tr td:nth-of-type(2)', // Latin Extended-B
  'table.wikitable:nth-of-type(7) tr td:nth-of-type(2)', // Latin Extended Additional
  'table.wikitable:nth-of-type(8) tr td:nth-of-type(2)', // IPA Extensions
  'table.wikitable:nth-of-type(9) tr td:nth-of-type(2)', // Spacing modifier letters
  'table.wikitable:nth-of-type(10) tr td:nth-of-type(2)', // Greek and Coptic
  'table.wikitable:nth-of-type(11) tr td', // Greek Extended
  'table.wikitable:nth-of-type(12) tr td:nth-of-type(2)', // Cyrillic
  'table.wikitable:nth-of-type(13) tr td:nth-of-type(2)', // Unicode symbols
  'table.wikitable:nth-of-type(14) tr td', // General Punctuation
  'table.wikitable:nth-of-type(15) tr td', // Superscripts and Subscripts
  'table.wikitable:nth-of-type(18) tr td', // Number Forms
  'table.wikitable:nth-of-type(19) tr td', // Arrows
  'table.wikitable:nth-of-type(20) tr td', // Mathematical symbols
  'table.wikitable:nth-of-type(21) tr td', // Miscellaneous Technical
  'table.wikitable:nth-of-type(22) tr td', // Enclosed Alphanumerics
  'table.wikitable:nth-of-type(23) tr td', // Box Drawing
  'table.wikitable:nth-of-type(24) tr td:nth-of-type(2)', // Block Elements
  'table.wikitable:nth-of-type(25) tr td:nth-of-type(2)', // Geometric Shapes
  'table.wikitable:nth-of-type(26) tr td', // Miscellaneous Symbols
  'table.wikitable:nth-of-type(28) tr td' // Alphabetic Presentation Forms
]

const candidates = []
let i = 1
for (const glyphSelector of glyphSelectors) {
  console.log('Selector:', glyphSelector)
  const glyphCells = document.querySelectorAll(glyphSelector)
  for (const cell of glyphCells) {
    const span = document.createElement('span')
    if (cell.children[0] && cell.children[0].tagName === 'A') {
      span.textContent = cell.children[0].textContent
    } else {
      span.textContent = cell.textContent
    }
    cell.replaceWith(span)
    if (
      span.offsetWidth > 0 &&
      span.offsetWidth <= 6 &&
      span.textContent.charCodeAt(0) !== 160 &&
      span.textContent.length <= 2 &&
      span.textContent.slice(0, 1) != '\\'
    ) {
      console.log(
        ` ${i++}: ${span.offsetWidth} ${span.textContent.trim()} ` +
          span.textContent.charCodeAt(0).toString(16)
      )
      candidates.push(span.textContent.trim())
    }
  }
  console.log('Candidates:', candidates.length, `"${candidates.join('')}"`)
  // Jim
  // const selected = "!'(),-./;I]`fijlrt|¡¦¨ª¯²³´·¸¹ºÌÍÎÏìíîïĨĩĪīĬĭĮįİıĵĺļłŕŗřţŧſƖƗƚƫƭǂǃǏǐǰȈȉȊȋȑȓțȶɉɍḟṫẛɼʰʱʲʳʴʵʶʸʹʻʼʽʾʿˀˁˆˇˈˉˊˋˌˍˎˏːˑ˒˓˔˕˖˗˘˙˚˛˜˝ˠˡˢˣˤ˪˫ˬ˭˯˰˱˲˳˴˵˶˷˸˹˺˻˼˿ʹ͵ͺ;Ϳ΄΅·ΐίιϊϳἰἱἲἳἴἵἶἷὶί᾽ι᾿῀῁῍῎῏ῐῑῒΐῖῗ῝῞῟῭΅`´῾ІЇіїј҂Ӏӏ‘’‚‛′‹›‾⁄‐‖‘’‚‛․‧′‵‹›‾⁃⁄⁅⁆⁏⁞ⁱ⁽⁾ⁿ₍₎ₐₑₒₓₔⅠⅰⅼ∕∙∣∫≀⋅⌈⌉⌊⌋⍘⍪⍮⍳⍸⍽⍿⎸⎹◜⚕".slice(
  // const selectedBase = "fijlrtªºÌÍÎÏìíîïĨĩĪīĬĭĮįİıĵĺļłŕŗřţŧſƖƗƚƫƭǂǃǏǐǰȈȉȊȋȑȓțȶɉɍḟṫẛɼʰʱʲʳʴʵʶʸʹʻʼʽʾʿˀˁˆˇˈˉˊˋˌˍˎˏːˑ˒˓˔˕˖˗ˠˡˢˣˤ˪˫ˬ˭˯˰˱˲˳˴˵˶˷˸˹˺˻˼˿ʹͺͿΐίιϊϳἰἱἲἳἴἵἶἷὶίιῐῑῒΐῖῗІЇіїјӀӏₐₑₒₓₔⅠⅰⅼ"
  const selectedBase =
    "!'(),-./;IJ]`fijlrst|¡¦¨ª¯°²³´·¸¹º" +
    'ÌÍÎÏìíîïĨĩĪīĬĭĮįİıĴĵĺļľŀłŕŗřśŝşšţŧſƖƗƚƫƭǁǃ' +
    'ǏǐǰȈȉȊȋȑȓșțɉḟṡṫẛʰʱʲʳʴʵʶʸʹʻʼʽʾʿˀˁ˂˃˄˅ˆˇˈˊˋ' +
    'ˌˍˎˏːˑ˒˓˘˙˚˛˜˝˞˟ˠˡˢˣˤ˥˦˧˨˩˪˫ˬ˯˰˱˲˳˴˷˸˹˺˻˼˽˾ͱʹ' +
    '͵ͺ;Ϳ΄΅·ΐΙΪίιϊϳἰἱἲἳἴἵἶἷὶί᾽ι᾿῀῁῍῎῏ῐῑῒΐῖῗῘῙ῝῞῟῭΅`´' +
    '῾ІЇЈгзѓѕіїјґғҙӀӟ‘’‚‛•′‹›⁄‘’‚‛•․‧′‵‸‹›⁃⁄⁅⁆⁏⁚⁝⁞⁰ⁱ⁵⁹⁽⁾' +
    'ⁿ₀₄₈₌₍₎⅟Ⅰⅰⅼ∘∣∥∫∶≀≬⊦⊧⋄〈〉⍿⎛⎜⎝⎞⎟⎠⎡⎣⎤⎥⎦' +
    'ꜣꝇꝉꝩꝲꝼꞁꞄꞅ꞉꞊Ꞌꞌ' + // Latin-D
    'ɨɩɭɺɽɿʇʈ' + // IPA
    'ᴉᴊᴬᴮᴯᴰᴱᴲᴳᴴᴵᴶᴷᴸᴹᴺᴻᴼᴽᴾᴿᵀᵁᵃᵄᵅᵇᵈᵉᵊᵋᵌᵍᵎᵏᵑᵒᵓᵔᵕᵖᵗᵘᵙᵛᵜᵝᵞᵟᵡᵢᵣᵤᵥᵦᵧᵨᵪᵮᵲᵳᵵᵻᵼ' + // phonetic
    "ﺍﻩﺁﺓ،؛؞؟ءآأؤإاةـهوٱٲٳۀۄۅۆۇۈۉۊۋۏ۔ەۣۥۦ" // arabic

  const selectedBaseSet = new Set(selectedBase.split(''))
  const reject =
    "!'(),-./;]`|¡¦¨ª¯²³´·¸¹°˘˙˚˛˜˝͵ͺ;΄΅·᾽ι᾿῀῁῍῎῏῝῞῟῭΅`´῾˦˧˨˩˪˫˥" +
    '‘’‚‛•′‹›⁄‘’‚‛•․‧′‵‸‹›⁃⁄⁅⁆⁏⁚⁝⁞⁰ⁱ⁵⁹⁽⁾ⁿ₀₄₈₌₍₎⅟Ⅰⅰⅼ∘∣∥∫∶≀≬⊦⊧⋄〈〉⍿⎛⎜⎝⎞⎟⎠⎡⎣⎤⎥⎦' +
    "،؛؞؟"
  const rejectSet = new Set(reject.split(''))
  const selectedSet = new Set()
  for (const glyph of selectedBaseSet) {
    if (!rejectSet.has(glyph)) {
      selectedSet.add(glyph)
    }
  }
  const selected = [...selectedSet].join('').slice(0, 256)
  console.log('Selected:', selected.length, `"${selected}"`)
}
