package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GAw9g4OYkACIgycDMXlienpqUX6UFovqzg/LzSElYFRuEU/4U2Ef/9tF5H934MjmI+rctq+y13KxeIRuvIIp+utlX8f+XTavjtTsXpLVCZ30H/d2w1e4TLT8nx3a20vUNDN4q6ReOT5LVqh/23v4/83VbUbH03+++ffz73335/1zou/YOLrHPO+iVVmbXa8VqNQWWvBhKqc9a5Z5hszmZxfS+26kbHQw+sNn+vrW29nLmWMST8i2HKr8d2z41+U23RvfOG/77NH9v/dPvsfcVXlTy+HTxbPP1b/7lSz7f+F2wufCv98cULeZvrv70u+/3k/d/1ijlPLFRaxn5p98+cBhu68ZTmTT1ltZb3mOuFqmdMe1h53y9obGa33xfa684nF2v2XPKv3pWrLS4Nbilrf3sytym86wXGVm2/X6Y4X568aV6bmPuLOtXhi/CJ/xyLpiO3FotuNL32LMuNuDzx8NTh27uUDfZeVrfh+XdlzbmVv1k/RGzofLX8+FK/8X3pjatW/HdmRM1iPOl12X//34L7/x5P/H4x7kfFnuvp7nsA+SeZ3AbOZz79tDNp5I0t035WiBi35PXOq7vvds5vmz1/kNG+G5q4HQde93teyzT3wvtqX2eqImqOM3D3nx14cq1cfuDvvRJjWvk17z00/553m7/nHvUHswOYDrnxnJ2b9D1Q4efFdyQ8VGfX/2h+OdJrbPb2yT8p2jnZVzHL3jRs56sIef5iiIff/X8GfdzPv7Dev+63Ff5SPry+UcY1F46RLXfsW84WU7pjKYNb7eVGJH09OztUinyn7+VKvqeoUzFRrltZLFFjgZyB+rqFJ9ONpladHniteCNtzfdqZ6v//7ibni+g5L1xcP/+vpkB74qEDnz5kWP3ZyXe4l+teG89+t95/TG/ipFoXHX3NzfemWDzi/5s3kb4zsl4yCad1Hzh71X9ty1Wh9B4bm1kyb/Sqtz9oErwdU7H7w/7bq43rmNbE8L3Tdr/Ae/T7L+MXPBERH/f7THqwf1f3/zpPlUn/9DWfXy9KsztSsfj6uma9IFOuFy8ZhJcsUOfIT2yMLNwf9PZD/J9n9cYaGS5Sv25WPUk4pvYrfLpzwHVvx38Ws/b8+XpFU2Rrim7w95uJOxPOCrec8uN/kmDcv7VjcrTmia3ib55WTjcqL5n19bB89eKTm61DH3bqHmAJ09+nlXT3R5EFR9i++TY/l2ziOBp5hSHkX0WkQGjo3esrAz+9vHp/YrDv1sdhu3ourX16kemct2IQ565FRat3Vyw5tvri0kUXrbb/F9nCqD23RPoRs8WGpT13gqp9ozRtKnc39KasXCyjn/NfmIHh//8Ab3aOiXKTPsxgZmColmBggGVxBowszo7I4vBcDdKNrCbAm5FJhBlRRCCbDCoiYGBbI4jEW2AgjMLuFAgQYPjv+JiZAYvDWNlA8kwMTAydDAwMiiwgHiAAAP//5QNzR8AEAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
