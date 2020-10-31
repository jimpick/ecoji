console.log('Hi Jim!')

const glyphSelectors = [
  'table.wikitable:nth-of-type(2) tr td'
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
    const selectedBase = "fijlrtªºÌÍÎÏìíîïĨĩĪīĬĭĮįİıĵĺļłŕŗřţŧſƖƗƚƫƭǂǃǏǐǰȈȉȊȋȑȓțȶɉɍḟṫẛɼʰʱʲʳʴʵʶʸʹʻʼʽʾʿˀˁˆˇˈˉˊˋˌˍˎˏːˑ˒˓˔˕˖˗ˠˡˢˣˤ˪˫ˬ˭˯˰˱˲˳˴˵˶˷˸˹˺˻˼˿ʹͺͿΐίιϊϳἰἱἲἳἴἵἶἷὶίιῐῑῒΐῖῗІЇіїјӀӏₐₑₒₓₔⅠⅰⅼ" + 
    "ꝇꝉꞁ꞉Ꞌꞌ"
  const selected = selectedBase.slice(
    0,
    256
  )
  console.log(
    'Selected:',
    selected.length,
    `"${selected}"`
  )
}
