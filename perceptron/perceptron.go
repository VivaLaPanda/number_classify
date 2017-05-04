package perceptron

import (
	feat "github.com/vivalapanda/number_classify/features"
)

type Perceptron struct {
	weights       []float64
	learningConst float64
}

var num_epochs = 300

func New(numFeatures int, learningConst float64) *Perceptron {
	p := &Perceptron{}
	p.learningConst = learningConst

	for i := 0; i < numFeatures; i++ {
		p.weights = append(p.weights, 0.0)
	}

	return p
}

func (p *Perceptron) Learn(imgArray [][]int, expectedValue int) *Perceptron {
	featureVector := getFeatures(imgArray)

	multiplyArray := []float64{}
	normalizedSum := 0.0

	for i, _ := range p.weights {
		multiplyArray = append(multiplyArray, featureVector[i]*p.weights[i])
		normalizedSum += featureVector[i] * p.weights[i]
	}

	var actualValue int
	if normalizedSum >= 0 {
		actualValue = 1
	} else {
		actualValue = -1
	}

	if actualValue != expectedValue {
		for i, _ := range p.weights {
			p.weights[i] = p.weights[i] + p.learningConst*float64(expectedValue-actualValue)*featureVector[i]
		}
	}

	return p
}

func (p *Perceptron) Predict(imgArray [][]int) int {
	featureVector := getFeatures(imgArray)
	multiplyArray := []float64{}
	normalizedSum := 0.0

	for i, _ := range p.weights {
		multiplyArray = append(multiplyArray, featureVector[i]*p.weights[i])
		normalizedSum += featureVector[i] * p.weights[i]
	}

	var actualValue int
	if normalizedSum >= 0 {
		actualValue = 1
	} else {
		actualValue = -1
	}

	return actualValue
}

func (p *Perceptron) Train(imgArrays [][][]int, expectedValues []int) *Perceptron {
	for epochIdx := 0; epochIdx < num_epochs; epochIdx++ {
		for i, imgArray := range imgArrays {
			p = p.Learn(imgArray, expectedValues[i])
		}
	}

	return p
}

func (p *Perceptron) Test(imgArrays [][][]int, expectedValues []int) float64 {
	successes := 0.0
	failures := 0.0

	for i, imgArray := range imgArrays {
		actual := p.Predict(imgArray)
		expected := expectedValues[i]

		if actual == expected {
			successes++
		} else {
			failures++
		}
	}

	return successes / (successes + failures)
}

func getFeatures(imgArray [][]int) (features []float64) {
	density := feat.Density(imgArray)
	minHIntercepts, maxHIntercepts := feat.HorizontalIntercepts(imgArray)
	hSymmetry := feat.HorizontalSymmetry(imgArray)
	minVIntercepts, maxVIntercepts := feat.VertIntercepts(imgArray)
	vSymmetry := feat.VertSymmetry(imgArray)

	return []float64{
		-1.0,
		density,
		float64(minHIntercepts),
		float64(maxHIntercepts),
		hSymmetry,
		float64(minVIntercepts),
		float64(maxVIntercepts),
		vSymmetry}
}
