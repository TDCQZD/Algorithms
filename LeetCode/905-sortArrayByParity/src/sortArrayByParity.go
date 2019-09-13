package src

func sortArrayByParity(A []int) []int {
    i,j := 0, len(A) - 1 
        for i < j {
            if (A[i] % 2) > (A[j] % 2) {
               A[i],A[j] = A[j],A[i]
            }

            if (A[i] % 2 == 0){
                i++
            } 
            if (A[j] % 2 == 1) {
              j--  
            }
        }

        return A
}